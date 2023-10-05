package main

import "fmt"

type Employee struct {
	ID           string
	FullName     string
	Type         string
	NationalCode int
	Hour         float64
}

const (
	BaseSallary   = 7500000
	ExtraHour     = 120000
	SalaryPerHour = 165000
	//	Shift         = 65000
)

func main() {
	employee1 := Employee{
		"3fsg546",
		"Hossein Erfani",
		"FullTime",
		2547896541,
		8,
	}
	employee1.CalculateSalary()
}
func (employee Employee) FullTimeSalary() (salary float64, name string) {
	salary = BaseSallary + (ExtraHour*employee.Hour)*1.4
	name = employee.FullName
	return
}
func (employee Employee) PartTimeSalary() (salary float64, name string) {
	salary = employee.Hour * SalaryPerHour
	name = employee.FullName
	return
}
func (employee Employee) CalculateSalary() {
	if employee.Type == "FullTime" {
		salary, name := employee.FullTimeSalary()
		fmt.Printf("salary is: %3f \n name is: %s \n", salary, name)
	} else if employee.Type == "PartTime" {
		salary, name := employee.PartTimeSalary()
		fmt.Printf("salary is: %3f \n name is: %s \n", salary, name)
	}
	//We had to add another statemnet if we had other type and it breaks open-close principle
}

/*
func Shift()  {

}
*/
