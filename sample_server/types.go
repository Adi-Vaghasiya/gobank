package main

import "fmt"

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewUser(firstname, lastname string) *User {
	fmt.Printf("Your Firstname is %v and Your Lastname is %v ", firstname, lastname)
	return &User{
		FirstName: firstname,
		LastName:  lastname,
	}
}
