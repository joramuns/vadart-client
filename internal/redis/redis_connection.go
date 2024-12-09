package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func NewRDB() *Connection {
	ctx := context.Background()

	redisAddr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	redisUser := os.Getenv("REDIS_USERNAME")
	redisPass := os.Getenv("REDIS_PASSWORD")

	log.Println("REDIS_ADDR:", redisAddr)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: redisUser,
		Password: redisPass,
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Error in Redis connection:", err)
		return nil
	}
	log.Println("Redis connected:", pong)

	return &Connection{conn: rdb}
}
