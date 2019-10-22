package ctx

import "context"

var Ctx context.Context

func InitContext() {
	Ctx = context.Background()
}

func SetApiKey(apiKey string) {
	Ctx = context.WithValue(Ctx, "api_key", apiKey)
}

func GetApiKey() interface{} {
	return Ctx.Value("api_key")
}

func SetServiceId(serviceId int) {
	Ctx = context.WithValue(Ctx, "service_id", serviceId)
}

func GetServiceId() interface{} {
	return Ctx.Value("service_id")
}

func SetUserId(userId int) {
	Ctx = context.WithValue(Ctx, "user_id", userId)
}

func GetUserId() interface{} {
	return Ctx.Value("user_id")
}

func SetUserUuid(userUuid string) {
	Ctx = context.WithValue(Ctx, "user_uuid", userUuid)
}

func GetUserUuid() interface{} {
	return Ctx.Value("user_uuid")
}
