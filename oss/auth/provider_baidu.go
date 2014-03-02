package auth


type BaiduProvider struct {
	BaiduAuth
}


func (bp BaiduProvider) Acl() string {
	return "x-bs-acl"
}

func (bp BaiduProvider) ObjectCopy() string {
	return "x-bs-copy-source"
}

func (bp BaiduProvider) ObjectCopyDrt() string {
	return "x-bs-copy-source-directive"
}

func (bp BaiduProvider) ObjectCopyDrtForReplace() string {
	return "replace"
}


func (bp BaiduProvider) MetaPrefix() string {
	return "x-bs-meta-"
}

