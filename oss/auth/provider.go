package auth

func Provide(provider string) Authorize {
	var authorize Authorize

	var a BaiduAuth
	a.Key = "mzx6uUfGhzNiidxNuRjaEmTc"
	a.Secret = []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l")
	a.Host = "bcs.duapp.com"
	a.Provider = provider

	authorize = a

	return authorize
}
