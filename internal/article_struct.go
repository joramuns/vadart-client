package internal

import "time"

type Article struct {
	Status      bool      `json:"status"`
	ArticleId   string    `json:"article_id"`
	MinPrice    int       `json:"min_price"`
	MaxPrice    int       `json:"max_price"`
	WorkerId    int       `json:"worker_id"`
	OrderId     string    `json:"order_id"`
	TimeBought  time.Time `json:"time_bought"`
	PriceBought int       `json:"price_bought"`
}
