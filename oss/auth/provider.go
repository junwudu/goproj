package auth


type HeaderField interface {
	Acl() string
	ObjectCopy() string
	ObjectCopyDrt() string
	ObjectCopyDrtForReplace() string
	MetaPrefix() string
}


type Provider interface {
	HeaderField
	Auth
}

func GetProvider(provider string) Provider {
	switch provider {
	default:
		return BaiduProvider{GetAuth(provider).(BaiduAuth)}
	case "baidu":
		return BaiduProvider{GetAuth(provider).(BaiduAuth)}
	}

}
