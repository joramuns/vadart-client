package application

import (
	"vadart_redis_client/pkg/redis"
)

type Application struct {
	RDB *redis.Connection
}

func NewApplication() *Application {
	return &Application{
		RDB: redis.NewRDB(),
	}
}
