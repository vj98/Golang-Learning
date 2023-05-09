package main

import "fmt"

func main() {
	whatWasSaid, otherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, otherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}