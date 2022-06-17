package main

import "fmt"

type employee struct {
	employeeId int
	person
	contactInfo
	startDate string
}

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	zipCode int
}

func (e employee) printEmployee() {
	fmt.Printf("Employee ID: %v\nFirstName: %v\nLastName: %v\nEmail: %v\nZipCode: %v\nStartDate: %v\n", e.employeeId, e.firstName, e.lastName, e.email, e.zipCode, e.startDate)

	fmt.Println("---------------------")
}

func (e *employee) updateFistName(newName string) {
	e.firstName = newName
}
