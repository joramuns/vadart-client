package handlers

import (
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
	"strconv"
	"strings"
)

func Add(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		text := c.Message().Text
		parts := strings.SplitN(text, " ", 4)

		if len(parts) == 2 {
			err := rdb.AddItem(parts[1], 0, 999999)
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
	}
}
