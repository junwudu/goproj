package auth


import (

)

type AccessInfo struct {
	Key string
	Secret string
}


type Info struct {
	AccessInfo
	Host string
	Provider Provider
}



type Parser interface {
	parseTo(info *Info) bool
	//parse() (*Info, bool)
}

type Provider string


func (info *Info) Parse(provider Provider) {
	var parser Parser

	var ps = []Parser{baiduParser(provider)}

	for _, v := range ps {
		parser = v
		if ok := parser.parseTo(info); ok {
			break
		}
	}
}


type baiduParser Provider
func (p baiduParser) parseTo(info *Info) bool {
	info.Provider = Provider(p)
	if p != "baidu" {
		return false
	}

	info.Key = "mzx6uUfGhzNiidxNuRjaEmTc"
	info.Secret = "TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l"
	info.Host = "http://bcs.duapp.com"

	return true
}

func NewInfo(provider Provider) Info {
	var info Info
	info.Parse(provider)
	return info
}

