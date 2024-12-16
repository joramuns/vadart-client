package handlers

import (
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
	"log"
)

func ClearAll(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := rdb.ClearAll()
		if err != nil {
			log.Println("Error refreshing:", err)
		}
		return nil
	}
}
