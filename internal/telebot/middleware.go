package telebot

import (
	tele "gopkg.in/telebot.v4"
	"log"
)

const owner = int64(89268804)

func Authorize() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Sender().ID == owner {
				return next(c)
			} else {
				log.Println("Unauthorized access denied:", c.Sender().ID, c.Sender().FirstName)
				return nil
			}
		}
	}
}
