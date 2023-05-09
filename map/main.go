package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// var myotherMap map[string]string // not able use this way
	myMap := make(map[string]User)

	me := User{
		FirstName: "Vijay",
		LastName:  "Bhati",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
}
