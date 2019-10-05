package ctx

import "context"

var Ctx context.Context

func InitContext() {
	Ctx = context.Background()
}

func SetApiKey(apiKey string) {
	Ctx = context.WithValue(Ctx, "Api-Key", apiKey)
}

func GetApiKey() interface{} {
	return Ctx.Value("Api-Key")
}

func SetToken(token string) {
	Ctx = context.WithValue(Ctx, "Token", token)
}

func GetToken() interface{} {
	return Ctx.Value("Token")
}
