package main

import "fmt"

func main() {
	var fullTime FullTimeEmployee
	fullTime.ID = "1"
	fullTime.FullName = "Reza"
	fullTime.NationalCode = 2547896541
	fullTime.ExtraHour = 30

	var partTime PartTimeEmployee
	partTime.ID = "2"
	partTime.FullName = "Negar"
	partTime.NationalCode = 5587496541
	partTime.PartTimeHour = 120

	var employee EmployeeInterface = partTime
	salary, name := employee.SalaryCalculator(partTime.PartTimeHour)
	fmt.Printf("salary is: %3f \n name is: %s \n", salary, name)
}

type EmployeeInterface interface {
	SalaryCalculator(hour float64) (salary float64, name string)
}

type FullTimeEmployee struct {
	ID           string
	FullName     string
	NationalCode int
	ExtraHour    float64
}

type PartTimeEmployee struct {
	ID           string
	FullName     string
	NationalCode int
	PartTimeHour float64
}

func (fullEmployee FullTimeEmployee) SalaryCalculator(hour float64) (salary float64, name string) {
	salary = 8000000 + (fullEmployee.ExtraHour)*1.4
	name = fullEmployee.FullName
	return
}

func (partEmployee PartTimeEmployee) SalaryCalculator(hour float64) (salary float64, name string) {
	salary = partEmployee.PartTimeHour * 320000
	name = partEmployee.FullName
	return
}
