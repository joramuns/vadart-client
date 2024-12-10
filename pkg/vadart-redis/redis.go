package vadart_redis

import (
	"github.com/go-redis/redis/v8"
)

type Connection struct {
	Conn *redis.Client
}
