package internal

import "github.com/go-redis/redis/v8"

type Application struct {
	RDB *redis.Client
}

func NewApplication() *Application {
	return &Application{
		RDB: GetRDB(),
	}
}
