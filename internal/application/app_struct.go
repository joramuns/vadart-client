package application

import (
	rdb "vadart_redis_client/internal/redis"
)

type Application struct {
	RDB *rdb.Connection
}

func NewApplication() *Application {
	return &Application{
		RDB: rdb.NewRDB(),
	}
}
