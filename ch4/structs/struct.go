package main

import (
	"fmt"
	"time"
)

func main() {
	// Structs
	type Employee struct {
		ID 		 int
		Name 	 string
		Address  string
		DoB 	 time.Time
		Position string
		Salary 	 int
		ManagerID int
	}	

	newEmployee := Employee {
		ID: 34, 
		Name: "Dilbert",
		Address: "Seattle",
		DoB:time.Date(2023, time.September, 12, 0, 0, 0, 0, time.UTC),
		Position: "Senior Haskell Programmer",
		Salary: 90000,
		ManagerID: 2,
	}
	//Print name 
	fmt.Println(newEmployee.Name)
	fmt.Println(newEmployee.DoB)
	pos := &newEmployee.Salary 
	
	//Promotion deets
	var employeeOfTheMonth *Employee = &newEmployee
	(*employeeOfTheMonth).Position += ""
	fmt.Print(*pos)
	str := "hello"
	fmt.Print(string(str[1]))
}
