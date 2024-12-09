package application

import (
	"github.com/joramuns/vadart-client/pkg/vadart-redis"
)

type Application struct {
	RDB *vadart_redis.Connection
}

func NewApplication() *Application {
	return &Application{
		RDB: vadart_redis.NewRDB(),
	}
}
