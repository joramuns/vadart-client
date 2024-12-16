package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/joramuns/vadart-client/internal/telebot"
	"github.com/joramuns/vadart-client/internal/telebot/handlers"
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

const owner = 89268804

func SubscribeTrading(rdb *vr.Connection, b *tele.Bot) {
	channelName := "articles"
	pubsub := rdb.Conn.Subscribe(context.Background(), channelName)
	channel := pubsub.Channel()

	for {
		select {
		case message := <-channel:
			if message.Payload[:3] == "del" {
				article := "Deleted: " + message.Payload[3:len(message.Payload)]
				msg, err := b.Send(tele.ChatID(owner), article)
				if err != nil {
					log.Println(err, "message:", msg)
				}
			} else if message.Payload[:3] == "add" {
				article := "Added: " + message.Payload[3:len(message.Payload)]
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
		Token:     os.Getenv("TOKEN"),
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/show", handlers.Show(rdb), telebot.Authorize())
	b.Handle("/refresh", handlers.Refresh(rdb), telebot.Authorize())
	b.Handle("/clearall", handlers.ClearAll(rdb), telebot.Authorize())
	b.Handle("/add", handlers.Add(rdb), telebot.Authorize())
	b.Handle("/sleep", handlers.Sleep(rdb), telebot.Authorize())

	go SubscribeTrading(rdb, b)

	b.Start()
}
