package main

import "log"

func main() {
	var isTrue bool

	isTrue = false

	if isTrue == true {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// Switch case

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")

	case "dog":
		log.Println("Cat is set to dog")

	case "fish":
		log.Println("Cat is set to fish")

	default:
		log.Println("cat is something else")
	}
}
