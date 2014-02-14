package auth


import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"bytes"
	"net/url"
)

type Auth struct {
	Info
	Method string
	Bucket string
	Object string
	Time string
	Ip string
	Size string
}


func NewAuth(provider Provider) Auth {
	info := NewInfo(provider)
	return Auth{info, "Get", "", "/", "", "", ""}
}


func (auth *Auth) SignedUrl() string {

	var u bytes.Buffer

	s := auth.Sign()

	u.WriteString(auth.Host)
	u.WriteString("/")
	u.WriteString(auth.Bucket)
	u.WriteString("/")
	u.WriteString(url.QueryEscape(auth.Object[1:]))
	u.WriteString("?sign=")
	u.WriteString(s)
	if auth.Time != "" {
		u.WriteString("&Time="); u.WriteString(auth.Time)
	}

	if auth.Ip != "" {
		u.WriteString("&Ip="); u.WriteString(auth.Ip)
	}

	if auth.Size != "" {
		u.WriteString("&Size="); u.WriteString(auth.Size)
	}

	return u.String()
}


func (auth *Auth) Sign() string {

	switch auth.Provider {
	case "baidu":
		flag, content := auth.contentBaidu()
		if flag == "" {
			panic("flag is empty!")
		}
		var buff bytes.Buffer
		buff.WriteString(flag); buff.WriteString(":")
		buff.WriteString(auth.Key); buff.WriteString(":")
		buff.WriteString(url.QueryEscape(auth.b6hms1(content)))
		return buff.String()

	default:
		return ""
	}
}

//base64 hmac sha1
func (auth *Auth) b6hms1(data []byte) string {
	h := hmac.New(sha1.New, []byte(auth.Secret))
	h.Write(data)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}


func (auth *Auth) contentBaidu() (string, []byte) {
	var flag bytes.Buffer
	var ct bytes.Buffer

	flag.WriteString("MBO")
	ct.WriteString("Method="); ct.WriteString(auth.Method); ct.WriteString("\n")
	ct.WriteString("Bucket="); ct.WriteString(auth.Bucket); ct.WriteString("\n")
	ct.WriteString("Object="); ct.WriteString(auth.Object); ct.WriteString("\n")

	if auth.Time != "" {
		flag.WriteString("T")
		ct.WriteString("Time="); ct.WriteString(auth.Time); ct.WriteString("\n")
	}

	if auth.Ip != "" {
		flag.WriteString("I")
		ct.WriteString("Ip="); ct.WriteString(auth.Ip); ct.WriteString("\n")
	}

	if auth.Size != "" {
		flag.WriteString("S")
		ct.WriteString("Size="); ct.WriteString(auth.Size); ct.WriteString("\n")
	}

	cts := ct.String()
	ct.Reset()
	ct.WriteString(flag.String()); ct.WriteString("\n"); ct.WriteString(cts)
	return flag.String(), ct.Bytes()
}
