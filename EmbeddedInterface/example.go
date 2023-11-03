package main

import "fmt"

type Animal interface {
	Sleep()
	Walk()
	Eat()
}

type Human interface {
	Animal //embedded interface
	Think()
	Speak()
}

type Employee struct {
	Name string
	Age  int
}

type Cat struct {
	Name string
}

func main() {
	cat := Cat{
		"Kitty",
	}
	employee := Employee{
		"Hana",
		29,
	}
	cat.Eat()
	cat.Walk()

	employee.Think()
	employee.Sleep()
}
func (cat *Cat) Sleep() {
	fmt.Println("animal is sleeping")
}

func (cat *Cat) Walk() {
	fmt.Println("animal is walking")
}

func (cat *Cat) Eat() {
	fmt.Println("animal is eating")
}

func (employee *Employee) Sleep() {
	fmt.Println("employee is sleeping")
}

func (employee *Employee) Walk() {
	fmt.Println("employee is walking")
}

func (employee *Employee) Eat() {
	fmt.Println("employee is eating")
}

func (employee *Employee) Think() {
	fmt.Println("employee is thinking")
}

func (employee *Employee) Speak() {
	fmt.Println("employee is speaking")
}
