package main

import (
	"fmt"
	"github.com/joho/godotenv"
	vadart_redis "github.com/joramuns/vadart-client/pkg/vadart-redis"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load(".env")
	rdb := vadart_redis.NewRDB()

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/show", func(c tele.Context) error {
		var message string
		articles := rdb.ShowAll()
		for _, article := range articles {
			message += fmt.Sprintf("ID: %s\nStatus: %t\n", article.ArticleId, article.Status)
			if !article.Status {
				message += fmt.Sprintf("Price: %d\nTime: %s\nBot: %s\n", article.PriceBought, article.TimeBought, article.OrderId)
			}
		}
		return c.Send(message)
	})

	b.Start()
}
