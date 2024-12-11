package application

import (
	"encoding/json"
	"fmt"
	"log"
)

func (a *Application) AddItem() {
	var (
		id       string
		minPrice int
		maxPrice int
	)
	fmt.Println("Enter article ID:")
	fmt.Scanln(&id)
	fmt.Println("Enter minimum price:")
	fmt.Scanln(&minPrice)
	fmt.Println("Enter maximum price:")
	fmt.Scanln(&maxPrice)
	err := a.RDB.AddItem(id, minPrice, maxPrice)
	if err != nil {
		log.Println("Error in add item application:", err)
	}
}

func (a *Application) UpdateItem() {
	var (
		id          string
		orderId     string
		priceBought int
	)
	fmt.Println("Enter article id:")
	fmt.Scanln(&id)

	fmt.Println("Enter order id:")
	fmt.Scanln(&orderId)
	fmt.Println("Enter price:")
	fmt.Scanln(&priceBought)
	err := a.RDB.UpdateItem(id, orderId, priceBought)
	if err != nil {
		log.Println("Error updating item:", err)
	}
}

func (a *Application) ClearID() {
	var id string
	fmt.Println("Enter ID:")
	fmt.Scanln(&id)
	err := a.RDB.ClearID(id)
	if err != nil {
		log.Println("Error clearing ID:", err)
	}
}

func (a *Application) ClearAll() {
	err := a.RDB.ClearAll()
	if err != nil {
		log.Println("ClearAll error:", err)
	}
}

func (a *Application) RefreshPubsub() {
	err := a.RDB.RefreshPubSub()
	if err != nil {
		log.Println("RefreshPubsub error:", err)
	}
}

func (a *Application) ShowAll() {
	articles := a.RDB.ShowAll()
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Println("error marshalling in Application ShowAll:", err)
	}
	fmt.Println(string(data))
}

func (a *Application) Command() {
	var (
		receiver string
		command  string
		value    string
	)
	fmt.Println("Enter bot name:")
	fmt.Scanln(&receiver)
	fmt.Println("Enter command:")
	fmt.Scanln(&command)
	fmt.Println("Enter value:")
	fmt.Scanln(&value)
	err := a.RDB.Command(receiver, command, value)
	if err != nil {
		log.Println("Error in command:", err)
	}
}
