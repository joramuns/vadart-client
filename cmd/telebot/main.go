package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	vadart_redis "github.com/joramuns/vadart-client/pkg/vadart-redis"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func SubscribeTrading(rdb *vadart_redis.Connection, b *tele.Bot) {
	channelName := "articles"
	pubsub := rdb.Conn.Subscribe(context.Background(), channelName)
	channel := pubsub.Channel()
	msg, err := b.Send(telebot.ChatID(89268804), "Subscribed to 'articles'")
	if err != nil {
		log.Fatal(err, "message:", msg)
	}
	for {
		select {
		case message := <-channel:
			if message.Payload[:3] == "del" {
				article := "Article bought:" + message.Payload[3:len(message.Payload)]
				msg, err := b.Send(telebot.ChatID(89268804), article)
				if err != nil {
					log.Fatal(err, "message:", msg)
				}
			}
		}
	}
}

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
	go SubscribeTrading(rdb, b)

	b.Start()
}
