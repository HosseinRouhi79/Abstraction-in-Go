package main

import "fmt"

func main() {

	person := struct {
		Name       string
		FamilyName string
		Age        int
		Score      float64
		Gender     string
	}{
		"Reza",
		"Mohammadi",
		24,
		14.4,
		"Male",
	}

	fmt.Printf("Person is: %+v\n", person)
}
