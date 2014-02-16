package provider

import (
	"github.com/junwudu/oss/auth"
	"crypto/sha1"
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"bytes"
)


func BaiduDo(auth *auth.Auth, signParam *auth.SignParameter) string {
	flag, data := content(signParam)

	h := hmac.New(sha1.New, auth.Secret)
	h.Write(data)
	signed := base64.URLEncoding.EncodeToString(h.Sum(nil))

	var buff bytes.Buffer
	buff.WriteString(flag); buff.WriteString(":")
	buff.WriteString(auth.Key); buff.WriteString(":")
	buff.WriteString(url.QueryEscape(signed))
	return buff.String()
}


func BaiduDo2Url(auth *auth.Auth, signParam *auth.SignParameter) string {

	var u bytes.Buffer

	s := BaiduDo(auth, signParam)

	u.WriteString(auth.Host)
	if auth.Host[len(auth.Host) - 1] != '/' {
		u.WriteString("/")
	}

	u.WriteString(signParam.Bucket)
	u.WriteString("/")

	if len(signParam.Object) > 0 {
		u.WriteString(url.QueryEscape(signParam.Object[1:]))
	}

	u.WriteString("?sign=")
	u.WriteString(s)
	if signParam.Time != "" {
		u.WriteString("&Time="); u.WriteString(signParam.Time)
	}

	if signParam.Ip != "" {
		u.WriteString("&Ip="); u.WriteString(signParam.Ip)
	}

	if signParam.Size != "" {
		u.WriteString("&Size="); u.WriteString(signParam.Size)
	}

	return u.String()
}


func content(signParam *auth.SignParameter) (string, []byte) {

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

