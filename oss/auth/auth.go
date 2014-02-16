package auth


import (
	"errors"
)

type AccessInfo struct {
	Key string
	Secret []byte
}


type Auth struct {
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



func (auth Auth) Sign(signPara *SignParameter) (result string, err error) {

	for _, processor := range AuthProcessors {
		if processor.Provider.MatchString(auth.Provider) {
			result = processor.Do(&auth, signPara)
		}
	}

	if len(result) < 1 {
		err = errors.New("provider not match")
	}

	return
}


func (auth Auth) SignedUrl(signPara *SignParameter) (result string, err error) {

	for _, processor := range AuthProcessors {
		if processor.Provider.MatchString(auth.Provider) {
			result = processor.Do2Url(&auth, signPara)
		}
	}

	if len(result) < 1 {
		err = errors.New("provider not match")
	}

	return
}




