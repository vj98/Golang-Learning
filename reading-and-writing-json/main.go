package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name": "Vijay",
			"last_name": "Bhati",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Rahul",
			"last_name": "Bhati",
			"hair_color": "black",
			"has_dog": true
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// write json from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Walley"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil {
		log.Println("error marshalling ", err)
	}

	fmt.Println(string(newJson))
}
