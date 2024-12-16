package handlers

import (
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
)

func Refresh(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := rdb.RefreshPubSub()
		if err != nil {
			return c.Send("Error refreshing:", err)
		}
		return c.Send("Refreshed!")
	}
}
