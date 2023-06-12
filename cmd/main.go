package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Title       string
	Description string
	Status      uint16
}

func main() {
	db, err := gorm.Open(mysql.Open("tickets.db"), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed")
	}

	fmt.Println("Connection successful")

	db.Create(&Ticket{Title: "Email not working", Description: "Help me please", Status: 1})

	var ticket Ticket
	db.First(&ticket, 1)
}
