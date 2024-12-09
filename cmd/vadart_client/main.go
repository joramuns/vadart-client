package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
	"vadart_redis_client/internal"
)

var articles = make(map[int]*internal.Article)

func showAll() {
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Println("Marshall error:", err)
	}
	fmt.Println(string(data))
}

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
	app := internal.NewApplication()

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
		} else if input == "show" {
			showAll()
		}
	}
}
