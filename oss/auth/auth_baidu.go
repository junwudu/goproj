package auth

import (
	"crypto/sha1"
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"bytes"
	"hash"
	"strings"
	"errors"
	"github.com/junwudu/goproj/oss/utils"
)


type BaiduAuth struct {
	AccessInfo
}



func (auth BaiduAuth) sign(signParam *SignParameter) (result string, err error) {
	flag, data := content(signParam)

	var f func() hash.Hash = sha1.New

	h := hmac.New(f, auth.Secret())
	h.Write(data)
	signed := base64.URLEncoding.EncodeToString(h.Sum(nil))

	var buff bytes.Buffer
	buff.WriteString(flag); buff.WriteString(":")
	buff.WriteString(auth.Key()); buff.WriteString(":")
	buff.WriteString(url.QueryEscape(utils.UrlPreEncode(signed)))
	result = buff.String()
	return
}


func (auth BaiduAuth) Sign(signParam *SignParameter) (sUrl string, err error) {

	var u bytes.Buffer

	s, err := auth.sign(signParam)

	u.WriteString("http://")

	if auth.Host() == "" {
		err = errors.New("Host is empty")
		return
	}

	u.WriteString(strings.Trim(auth.Host(), "/"))
	u.WriteString("/")

	u.WriteString(strings.Trim(signParam.Bucket, "/"))

	objStr := strings.TrimLeft(signParam.Object, "/")

	if len(objStr) > 0 {
		u.WriteString("/")
		u.WriteString(url.QueryEscape(utils.UrlPreEncode(objStr)))
	}

	u.WriteString("?sign=")
	u.WriteString(s)
	if signParam.Time != "" {
		u.WriteString("&time="); u.WriteString(signParam.Time)
	}

	if signParam.Ip != "" {
		u.WriteString("&ip="); u.WriteString(signParam.Ip)
	}

	if signParam.Size != "" {
		u.WriteString("&size="); u.WriteString(signParam.Size)
	}

	sUrl = u.String()
	return
}


func content(signParam *SignParameter) (string, []byte) {

	var flag bytes.Buffer
	var ct bytes.Buffer

	flag.WriteString("MBO")
	ct.WriteString("Method="); ct.WriteString(signParam.Method); ct.WriteString("\n")
	ct.WriteString("Bucket="); ct.WriteString(signParam.Bucket); ct.WriteString("\n")
	ct.WriteString("Object="); ct.WriteString(signParam.Object); ct.WriteString("\n")

	if signParam.Time != "" {
		flag.WriteString("T")
		ct.WriteString("Time="); ct.WriteString(signParam.Time); ct.WriteString("\n")
	}

	if signParam.Ip != "" {
		flag.WriteString("I")
		ct.WriteString("Ip="); ct.WriteString(signParam.Ip); ct.WriteString("\n")
	}

	if signParam.Size != "" {
		flag.WriteString("S")
		ct.WriteString("Size="); ct.WriteString(signParam.Size); ct.WriteString("\n")
	}

	cts := ct.String()
	ct.Reset()
	ct.WriteString(flag.String()); ct.WriteString("\n"); ct.WriteString(cts)

	return flag.String(), ct.Bytes()
}
