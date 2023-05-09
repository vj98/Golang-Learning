package main

import "log"

func main() {
	var mySlice []string // slice

	mySlice = append(mySlice, "Vijay")
	mySlice = append(mySlice, "Bhati")

	log.Println(mySlice)
}
