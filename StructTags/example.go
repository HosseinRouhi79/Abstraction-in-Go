package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID     int
	Name   string `json:"name"` //It changes Name to name in json result
	Age    int
	Height float64
	Gender string
}

func main() {
	person := Person{2, "Parvaneh", 24, 175, "Female"}

	personJson, _ := json.MarshalIndent(person, "", "   ")
	fmt.Println(string(personJson))
}
