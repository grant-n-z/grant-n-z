package ctx

import (
	"context"
	"github.com/google/uuid"
)

// Global context
var Ctx context.Context

// Initialize context
func InitContext() {
	Ctx = context.Background()
}

// Setter api key
func SetApiKey(apiKey string) {
	Ctx = context.WithValue(Ctx, "api_key", apiKey)
}

// Getter api key
func GetApiKey() interface{} {
	return Ctx.Value("api_key")
}

// Setter service id
func SetServiceId(serviceId int) {
	Ctx = context.WithValue(Ctx, "service_id", serviceId)
}

// Getter service id
func GetServiceId() interface{} {
	return Ctx.Value("service_id")
}

// Setter user id
func SetUserId(userId int) {
	Ctx = context.WithValue(Ctx, "user_id", userId)
}

// Getter user id
func GetUserId() interface{} {
	return Ctx.Value("user_id")
}

// Setter user uuid
func SetUserUuid(userUuid uuid.UUID) {
	Ctx = context.WithValue(Ctx, "user_uuid", userUuid)
}

// Getter user uuid
func GetUserUuid() interface{} {
	return Ctx.Value("user_uuid")
}
