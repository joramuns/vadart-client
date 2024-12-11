package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
)

const owner = 89268804

func SubscribeTrading(rdb *vr.Connection, b *tele.Bot) {
	channelName := "articles"
	pubsub := rdb.Conn.Subscribe(context.Background(), channelName)
	channel := pubsub.Channel()
	msg, err := b.Send(tele.ChatID(89268804), "Subscribed to 'articles'")
	if err != nil {
		log.Fatal(err, "message:", msg)
	}
	for {
		select {
		case message := <-channel:
			if message.Payload[:3] == "del" {
				article := "Article bought:" + message.Payload[3:len(message.Payload)]
				msg, err := b.Send(tele.ChatID(owner), article)
				if err != nil {
					log.Println(err, "message:", msg)
				}
			} else if message.Payload[:3] == "add" {
				article := "Start hunting article" + message.Payload[3:len(message.Payload)]
				msg, err := b.Send(tele.ChatID(owner), article)
				if err != nil {
					log.Println(err, "message:", msg)
				}
			}
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	rdb := vr.NewRDB()

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

	b.Handle("/refresh", func(c tele.Context) error {
		err := rdb.RefreshPubSub()
		if err != nil {
			c.Send("Error refreshing:", err)
		}
		return nil
	})

	b.Handle("/clearall", func(c tele.Context) error {
		err := rdb.ClearAll()
		if err != nil {
			log.Println("Error refreshing:", err)
		}
		return nil
	})

	b.Handle("/add", func(c tele.Context) error {
		text := c.Message().Text
		parts := strings.SplitN(text, " ", 4)

		if len(parts) == 2 {
			err = rdb.AddItem(parts[1], 0, 999999)
			if err != nil {
				return c.Send("Error in add item:", err)
			}
			return nil
		} else if len(parts) < 4 {
			return c.Send("Not enough arguments provided")
		}

		minPrice, err := strconv.Atoi(parts[2])
		if err != nil {
			return c.Send("Wrong min price argument")
		}
		maxPrice, err := strconv.Atoi(parts[3])
		if err != nil {
			return c.Send("Wrong max price argument")
		}

		err = rdb.AddItem(parts[1], minPrice, maxPrice)
		if err != nil {
			return c.Send("Error in add item:", err)
		}
		return nil
	})

	b.Handle("/sleep", func(c tele.Context) error {
		text := c.Message().Text
		parts := strings.SplitN(text, " ", 3)

		if len(parts) < 3 {
			return c.Send("Wrong sleep format - /sleep receiver seconds")
		}

		err := rdb.Command(parts[1], "sleep", parts[2])
		if err != nil {
			c.Send("Error in command sleep:", err)
		}
		return nil
	})

	go SubscribeTrading(rdb, b)

	b.Start()
}
