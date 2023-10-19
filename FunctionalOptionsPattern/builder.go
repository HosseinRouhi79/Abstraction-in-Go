package main

import "fmt"

type Person struct {
	ID     int
	Name   string
	Age    int
	Height float64
	Gender string
}

type PersonBuilder struct {
	Person
}

func (builder *PersonBuilder) SetID(id int) *PersonBuilder {
	builder.ID = id
	return builder
}

func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}

func (builder *PersonBuilder) SetAge(age int) *PersonBuilder {
	builder.Age = age
	return builder
}

func (builder *PersonBuilder) SetHeight(height float64) *PersonBuilder {
	builder.Height = height
	return builder
}

func (builder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	builder.Gender = gender
	return builder
}

func (builder *PersonBuilder) Build() Person {
	return builder.Person
}

func main() {
	person1 := &PersonBuilder{}
	person1.SetName("Reza").SetGender("male").SetAge(22).SetID(12).SetHeight(172.2).Build()
	fmt.Printf("Person : %+v\n", *person1)
}
