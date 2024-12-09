package redis

import (
	"github.com/go-redis/redis/v8"
)

type Connection struct {
	conn *redis.Client
}
