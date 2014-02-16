package auth

func Provide(provider string) {
	var authorize auth.Auth
	authorize.Provider = provider
}
