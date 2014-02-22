package auth

type AccessInfo struct {
	Key string
	Secret []byte
}


type AuthInfo struct {
	AccessInfo
	Host string
	Provider string
}


type SignParameter struct {
	Method string
	Bucket string
	Object string
	Time string
	Ip string
	Size string
}


type Authorize interface {
	Sign(*SignParameter) (string, error)
}
