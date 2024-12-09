package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
	"vadart_redis_client/internal"
	"vadart_redis_client/internal/application"
)

var articles = make(map[int]*internal.Article)

func addItem() {
	var id int
	fmt.Println("Enter article ID:")
	fmt.Scanln(&id)
	articles[id] = &internal.Article{
		Status:    true,
		ArticleId: id,
	}
	fmt.Println("Enter minimum price:")
	fmt.Scanln(&articles[id].MinPrice)
	fmt.Println("Enter maximum price:")
	fmt.Scanln(&articles[id].MaxPrice)
}

func delItem() {
	var id int
	fmt.Println("Enter article id:")
	fmt.Scanln(&id)
	article, exists := articles[id]
	if exists {
		article.Status = false
		fmt.Println("Enter order id:")
		fmt.Scanln(&article.OrderId)
		article.TimeBought = time.Now()
		fmt.Println("Enter price:")
		fmt.Scanln(&article.PriceBought)
	} else {
		fmt.Println("ID not found")
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	app := application.NewApplication()

	var input string
	for {
		fmt.Println("Main menu:")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		} else if input == "add" {
			addItem()
		} else if input == "del" {
			delItem()
		} else if input == "clear" {
			var id string
			fmt.Println("Enter ID:")
			fmt.Scanln(&id)
			app.ClearID(id)
		} else if input == "show" {
			app.ShowAll()
		}
	}
}
