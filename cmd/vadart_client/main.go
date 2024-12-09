package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"vadart-client/internal/application"
)

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
			app.AddItem()
		} else if input == "del" {
			app.UpdateItem()
		} else if input == "clear" {
			app.ClearID()
		} else if input == "show" {
			app.ShowAll()
		}
	}
}
