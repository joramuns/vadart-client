package handlers

import (
	"fmt"
	vr "github.com/joramuns/vadart-client/pkg/vadart-redis"
	tele "gopkg.in/telebot.v4"
)

func Show(rdb *vr.Connection) tele.HandlerFunc {
	return func(c tele.Context) error {
		var message string
		articles := rdb.ShowAll()
		if len(articles) == 0 {
			return c.Send("Empty")
		}
		for _, article := range articles {
			message += fmt.Sprintf("*ID: %s*\n", article.ArticleId)
			if article.Status {
				message += "Status: in progress\n"
				message += fmt.Sprintf("Max price: %d\n", article.MaxPrice)
			} else {
				message += "Status: done\n"
				message += fmt.Sprintf("Price: %d\nTime: %s\nInfo: %s\n", article.PriceBought, article.TimeBought.Format("15:04:05"), article.OrderId)
			}
			message += "===========\n"
		}
		return c.Send(message)
	}
}
