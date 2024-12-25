package handlers

import (
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
)

func Check(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := rdb.Command("ALL", "check", "")
		if err != nil {
			return c.Send("Error in command check:", err)
		}
		return nil
	}
}
