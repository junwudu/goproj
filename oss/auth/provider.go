package auth

func Provide(provider string) Auth {
	var authorize Auth
	authorize.Key = "mzx6uUfGhzNiidxNuRjaEmTc"
	authorize.Secret = []byte("TNmHdLcEPBUI1cmNr2GD1tL5YRTvb72l")
	authorize.Host = "bcs.duapp.com"
	authorize.Provider = provider
	return authorize
}
