package main

import "fmt"

type Person2 struct {
	ID     int
	Name   string
	Age    int
	Height float64
	Gender string
}

type PersonOptions func(person *Person2)

func main() {
	NewPerson(SetName("Reza"), SetAge(22), SetHeight(180), SetGender("Male"))
}

func NewPerson(options ...PersonOptions) {
	person := Person2{}
	for _, option := range options {
		option(&person)
	}
	fmt.Printf("Person : %+v", person)
}

func SetName(name string) PersonOptions {
	return func(person *Person2) {
		person.Name = name
	}
}

func SetAge(age int) PersonOptions {
	return func(person *Person2) {
		person.Age = age
	}
}

func SetHeight(height float64) PersonOptions {
	return func(person *Person2) {
		person.Height = height
	}
}

func SetGender(gender string) PersonOptions {
	return func(person *Person2) {
		person.Gender = gender
	}
}
