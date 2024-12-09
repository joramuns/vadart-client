package application

import (
	"vadart_redis_client/pkg/vadart-redis"
)

type Application struct {
	RDB *vadart_redis.Connection
}

func NewApplication() *Application {
	return &Application{
		RDB: vadart_redis.NewRDB(),
	}
}
