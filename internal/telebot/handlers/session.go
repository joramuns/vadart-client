package handlers

import (
	"context"
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
	"strings"
)

func Session(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		text := c.Message().Text
		parts := strings.SplitN(text, " ", 2)
		if len(parts) < 2 {
			return c.Send("Wrong session format - /session session-id")
		}

		rdb.Conn.Set(context.Background(), "PHPSESSID", parts[1], -1)
		err := rdb.Command("ALL", "session", parts[1])
		if err != nil {
			return c.Send("Error in command check:", err)
		}
		return nil
	}
}
