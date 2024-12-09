package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func (c *Connection) ShowAll() {
	articles, err := c.conn.HGetAll(context.Background(), "articles").Result()
	if err != nil {
		log.Println("HGetAll in Redis error:", err)
	}
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Println("Marshall error:", err)
	}
	fmt.Println(string(data))
}

func (c *Connection) ClearID(id string) error {
	exists, err := c.conn.HExists(context.Background(), "articles", id).Result()
	if err != nil {
		return fmt.Errorf("hexists in Redis error: %v", err)
	}
	if !exists {
		return fmt.Errorf("field %s does not exist in hash 'articles'")
	}

	err = c.conn.HDel(context.Background(), "articles", id).Err()
	if err != nil {
		return fmt.Errorf("clear by ID error: %v", err)
	}
	return nil
}
