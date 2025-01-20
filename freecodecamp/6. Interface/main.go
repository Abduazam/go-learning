package main

import "fmt"

type Employee interface {
	getName() string
	getSalary() int
}

type Fulltime struct {
	Name   string
	Salary int
}

func (f Fulltime) getName() string {
	return f.Name
}

func (f Fulltime) getSalary() int {
	return f.Salary
}

type Contractor struct {
	Name      string
	HourlyPay int
	HoursYear int
}

func (c Contractor) getName() string {
	return c.Name
}

func (c Contractor) getSalary() int {
	return c.HourlyPay * c.HoursYear
}

func main() {
	ftEmployee := Fulltime{
		Name:   "Bob",
		Salary: 75000,
	}

	fmt.Printf("%s is full-time employee and salary is $%d\n", ftEmployee.getName(), ftEmployee.getSalary())

	cEmployee := Contractor{
		Name:      "Jane",
		HourlyPay: 50,
		HoursYear: 1920,
	}

	fmt.Printf("%s is controctor employee and salary is $%d\n", cEmployee.getName(), cEmployee.getSalary())
}
