package auth

type AccessInfo struct {
	key string
	secret []byte
	host string
	provider string
}

func (ac AccessInfo) Name() string {
	return ac.provider
}

func (ac AccessInfo) Key() string {
	return ac.key
}

func (ac AccessInfo) Host() string {
	return ac.host
}

func (ac AccessInfo) Secret() []byte {
	return ac.secret
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


type Auth interface {
	Name() string
	Host() string
	Key() string
	Secret() []byte
	Authorize
}


func GetAuth (ps string) Auth {
	switch ps {
	default:
		return BaiduAuth{AccessInfo{"mzx6uUfGhzNiidxNuRjaEmTc", []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l"), "bcs.duapp.com", ps}}
	case "baidu":
		return BaiduAuth{AccessInfo{"mzx6uUfGhzNiidxNuRjaEmTc", []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l"), "bcs.duapp.com", ps}}
	}
}
