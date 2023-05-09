package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         string
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "11111113232",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
