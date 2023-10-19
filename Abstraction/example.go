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

	var employee EmployeeSalaryCalculator = fullTime
	salary, name := employee.SalaryCalculate(fullTime.ExtraHour)
	fmt.Printf("salary is: %3f \n name is: %s \n", salary, name)
}

type EmployeeSalaryCalculator interface {
	SalaryCalculate(hour float64) (salary float64, name string)
}

type BaseEmployee struct {
	ID           string
	FullName     string
	NationalCode int
}

type FullTimeEmployee struct {
	BaseEmployee
	ExtraHour float64
}

type PartTimeEmployee struct {
	BaseEmployee
	PartTimeHour float64
}

func (fullEmployee FullTimeEmployee) SalaryCalculate(hour float64) (salary float64, name string) {
	salary = 8000000 + (fullEmployee.ExtraHour)*1.4
	name = fullEmployee.FullName
	return
}

func (partEmployee PartTimeEmployee) SalaryCalculate(hour float64) (salary float64, name string) {
	salary = partEmployee.PartTimeHour * 320000
	name = partEmployee.FullName
	return
}
