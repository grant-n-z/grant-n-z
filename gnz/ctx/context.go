package ctx

import (
	"context"

	"github.com/google/uuid"
)

// Global context
var ctx context.Context

// Initialize context
func InitContext() {
	ctx = context.Background()
}

// Get ctx
func GetCtx() context.Context {
	return ctx
}

// Setter api key
func SetApiKey(apiKey string) {
	ctx = context.WithValue(ctx, "api_key", apiKey)
}

// Getter api key
func GetApiKey() interface{} {
	return ctx.Value("api_key")
}

// Setter service id
func SetServiceId(serviceId int) {
	ctx = context.WithValue(ctx, "service_id", serviceId)
}

// Getter service id
func GetServiceId() interface{} {
	return ctx.Value("service_id")
}

// Setter user id
func SetUserId(userId int) {
	ctx = context.WithValue(ctx, "user_id", userId)
}

// Getter user id
func GetUserId() interface{} {
	return ctx.Value("user_id")
}

// Setter user uuid
func SetUserUuid(userUuid uuid.UUID) {
	ctx = context.WithValue(ctx, "user_uuid", userUuid)
}

// Getter user uuid
func GetUserUuid() interface{} {
	return ctx.Value("user_uuid")
}
