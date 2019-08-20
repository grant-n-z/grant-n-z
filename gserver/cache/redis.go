package cache

type RedisClient interface {
	Get()

	Set()
}
