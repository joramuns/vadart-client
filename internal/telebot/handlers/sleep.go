package handlers

import (
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
	"strings"
)

func Sleep(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		text := c.Message().Text
		parts := strings.SplitN(text, " ", 3)

		if len(parts) < 3 {
			return c.Send("Wrong sleep format - /sleep receiver seconds")
		}

		err := rdb.Command(parts[1], "sleep", parts[2])
		if err != nil {
			return c.Send("Error in command sleep:", err)
		}
		return nil
	}
}
