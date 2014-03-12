package oss

import (
	"time"
	"net/http"
	"github.com/junwudu/goproj/oss/errors"
	"strings"
	"strconv"
	"bytes"
	"mime"
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"io"
	"crypto/md5"
	"encoding/hex"
)

/*oss object*/
type Object struct {
	Client *Client

	/*object name starting with '/'*/
	Name string

	/*alias of Name, used as download name*/
	Alias string

	/*start with '/', without bucket part */
	ParentDir string

	IsDir bool

	/*modify time */
	ModifyTime time.Time

	/*bucket that in */
	Bucket *Bucket

	/*starting pos in this bucket*/
	Pos uint64

	/*content type (common this field in http header)*/
	Type string

	Data []byte

	/*Range start*/
	Start uint64

	/*size of object by byte*/
	Size uint64

	MD5 string

	/*user defined info*/
	Meta map[string]string

	/*Acl of this object*/
	Acl string

	/*location*/
	Location string
}


func (object *Object) setName(name string) {
	if name[0] != '/' {
		object.Name = "/" + name
	} else {
		object.Name = name
	}

	object.Alias = object.Name[strings.LastIndex(object.Name, "/") + 1 : len(object.Name)]

	if object.Type == "" {
		rIdx := strings.LastIndex(object.Name, ".")
		if rIdx >= 0 {
			ext := object.Name[rIdx : len(object.Name)]
			object.Type = mime.TypeByExtension(ext)
		}
	}
}


func (object *Object) setData(data []byte) {
	object.Data = data
	object.Size = uint64(len(object.Data))
}


func (object *Object) available() (ok bool, err error) {

	if len (object.Name) < 2 || object.Name[0] != '/'  {
		err = errors.Error("object name is valid fail: " + object.Name)
	}

	if object.Bucket.Name == "" {
		err = errors.Error("bukcet is not set")
	}

	if object.Data == nil {
		err = errors.Error("not data")
	}

	if object.Type == "" {
		err = errors.Error("content-type is not set")
	}

	if object.Size == 0 {
		err = errors.Error("size is 0")
	}

	ok = err == nil
	return
}



func (object *Object) Delete() (err error) {
	url, err := object.Client.SignedUrl("DELETE", object.Bucket.Name, object.Name, "", "", "")

	if err != nil {
		return
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if (err != nil) {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, object.Client.Provider)

	return
}


func (object *Object) Put() (err error) {
	ok, err := object.available()
	if !ok {
		return
	}

	url, err := object.Client.SignedUrl("PUT", object.Bucket.Name, object.Name, "", "", "")
	if err != nil {
		return
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(object.Data))

	if err != nil {
		return
	}

	if req.Header.Get("Content-Length") == "" {
		req.Header.Set("Content-Length", strconv.FormatUint(object.Size, 10))
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", object.Type)
	}

	if req.Header.Get("Content-Disposition") == "" {
		req.Header.Set("Content-Disposition", object.Alias)
	}

	if req.Header.Get(object.Client.Provider.Acl()) == "" {
		req.Header.Set(object.Client.Provider.Acl(), object.Acl)
	}


	metaPrefix := object.Client.Provider.MetaPrefix()
	if len(object.Meta) > 0 {
		for k, v := range object.Meta {
			req.Header.Set(metaPrefix + k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, object.Client.Provider)

	if err == nil {
		md5 := resp.Header.Get("Content-MD5")
		if md5 == "" {
			err = errors.Error("put error! content md5 empty")
		} else {
			object.MD5 = md5
		}
	}

	return
}


func (object *Object) Get() (err error) {
	url, err := object.Client.SignedUrl("GET", object.Bucket.Name, object.Name, "", "", "")
	if err != nil {
		return
	}

	//download to file
	if object.Location != "" {
		fName := filepath.Join(object.Location, object.Alias)
		fp, err := os.OpenFile(fName, os.O_APPEND| os.O_CREATE, os.FileMode(0666))
		if err != nil {
			goto MEM
		}
		defer fp.Close()

		stat, err := fp.Stat()
		if err != nil {
			goto MEM
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", stat.Size()))
		req.Header.Set("response-content-disposition", object.Alias)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		object.Size = uint64(stat.Size())
		object.Start = uint64(stat.Size())
		wCount, err := io.Copy(fp, resp.Body)
		if err != nil {
			object.Size += uint64(wCount)
		}

		object.Type = resp.Header.Get("Content-Type")

		//modify time
		object.ModifyTime, _ = time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))

		object.MD5 = resp.Header.Get("ETag")

		return nil
	}

MEM:
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	object.Data, err = ioutil.ReadAll(resp.Body)
	object.Size = uint64(len(object.Data))

	object.Type = resp.Header.Get("Content-Type")

	//modify time
	object.ModifyTime, _ = time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))

	object.MD5 = resp.Header.Get("ETag")

	md := md5.Sum(object.Data)
	if mds := hex.EncodeToString(md[0:]); mds != object.MD5 {
		err = errors.Error(fmt.Sprintf("data md5:%s != %s(recived md5)", mds, object.MD5))
	}

	return
}


func (dstObject *Object) Copy(srcObject *Object, copyMeta bool) (err error) {
	url, err := dstObject.Client.SignedUrl("PUT", dstObject.Bucket.Name, dstObject.Name, "", "", "")
	if err != nil {
		return
	}

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return
	}

	req.Header.Set(dstObject.Client.Provider.ObjectCopy(), dstObject.Client.ObjectUrl(srcObject))
	if !copyMeta {
		req.Header.Set(dstObject.Client.Provider.ObjectCopyDrt(), dstObject.Client.Provider.ObjectCopyDrtForReplace())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, dstObject.Client.Provider)

	if err == nil {
		md5 := resp.Header.Get("Content-MD5")
		if md5 == "" {
			err = errors.Error("copy error! content md5 empty")
		} else {
			dstObject.MD5 = md5
		}
	}

	return

}


func (object *Object) Head() (err error) {
	url, err := object.Client.SignedUrl("HEAD", object.Bucket.Name, object.Name, "", "", "")
	if err != nil {
		return
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = errors.GetError(resp, object.Client.Provider)

	if err == nil {
		object.MD5 = resp.Header.Get("Content-MD5")
		object.Size, _ = strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
		object.Type = resp.Header.Get("Content-Type")

		//modify time
		object.ModifyTime, _ = time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))

		//meta headers
		metaPrefix := object.Client.Provider.MetaPrefix()
		for k, v := range resp.Header {
			if idx := strings.Index(k, metaPrefix); idx != -1 {
				if object.Meta == nil {
					object.Meta = make(map[string]string)
				}
				//just keep name without prefix
				object.Meta[k[idx+len(metaPrefix): len(k)]] = strings.Join(v, ";")
			}
		}
	}
	return
}


