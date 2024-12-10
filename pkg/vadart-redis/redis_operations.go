package vadart_redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (c *Connection) ShowAll() map[string]Article {
	articles, err := c.conn.HGetAll(context.Background(), "articles").Result()
	if err != nil {
		log.Println("HGetAll in Redis error:", err)
	}
	var articleMap = make(map[string]Article)
	for key, value := range articles {
		article, err := c.unmarshalItem([]byte(value))
		if err != nil {
			fmt.Println("Error in Connection ShowAll:", err)
		}
		articleMap[key] = article
	}

	return articleMap
}

func (c *Connection) ClearID(id string) error {
	exists, err := c.conn.HExists(context.Background(), "articles", id).Result()
	if err != nil {
		return fmt.Errorf("hexists in Redis error: %v", err)
	}
	if !exists {
		return fmt.Errorf("field %s does not exist in hash 'articles'", id)
	}

	err = c.conn.HDel(context.Background(), "articles", id).Err()
	if err != nil {
		return fmt.Errorf("clear by ID error: %v", err)
	}
	return nil
}

func (c *Connection) AddItem(id string, minPrice, maxPrice int) error {
	article := Article{
		Status:    true,
		ArticleId: id,
		MinPrice:  minPrice,
		MaxPrice:  maxPrice,
	}
	jsonData, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("error marshalling in AddItem: %v", err)
	}

	err = c.conn.HSet(context.Background(), "articles", id, jsonData).Err()
	if err != nil {
		return fmt.Errorf("error in Connection AddItem: %v", err)
	}

	c.conn.Publish(context.Background(), "articles", "add"+id)

	return nil
}

func (c *Connection) UpdateItem(id, orderId string, price int) error {
	data, err := c.conn.HGet(context.Background(), "articles", id).Result()
	if err != nil {
		return fmt.Errorf("get item error in del item: %v", err)
	}
	article, err := c.unmarshalItem([]byte(data))
	if err != nil {
		return fmt.Errorf("error in unmarshalling: %v", err)
	}

	article.TimeBought = time.Now()
	article.Status = false
	article.OrderId = orderId
	article.PriceBought = price

	err = c.pushItem(id, article)
	if err != nil {
		return fmt.Errorf("error in Connection UpdateItem: %v", err)
	}

	c.conn.Publish(context.Background(), "articles", "del"+id)

	return nil
}

func (c *Connection) CheckStatus(id string) bool {
	data, err := c.conn.HGet(context.Background(), "articles", id).Result()
	if err != nil {
		log.Println("get item error in check status:", err)
		return false
	}
	article, err := c.unmarshalItem([]byte(data))
	if err != nil {
		log.Println("error in unmarshalling:", err)
		return false
	}

	return article.Status
}

func (c *Connection) Command(receiver, command, value string) error {
	com := Settings{
		Receiver: receiver,
		Command:  command,
		Value:    value,
	}
	jsonData, err := json.Marshal(com)
	if err != nil {
		return fmt.Errorf("error marshalling in Command: %v", err)
	}
	err = c.conn.Publish(context.Background(), "master", jsonData).Err()
	if err != nil {
		return fmt.Errorf("error publishing in master: %v", err)
	}
	return nil
}

func (c *Connection) pushItem(id string, article Article) error {
	jsonData, err := json.Marshal(article)
	if err != nil {
		return fmt.Errorf("error marshalling in AddItem: %v", err)
	}

	err = c.conn.HSet(context.Background(), "articles", id, jsonData).Err()
	if err != nil {
		return fmt.Errorf("error in Connection AddItem: %v", err)
	}
	return nil
}

func (c *Connection) unmarshalItem(value []byte) (Article, error) {
	var article Article
	err := json.Unmarshal(value, &article)
	if err != nil {
		return Article{}, fmt.Errorf("unmarshall error in Connection GetItem: %v", err)
	}
	return article, nil
}
