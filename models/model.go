package models

type User struct {
	ID    int
	Name  string
	Email string
}

// DB
var Users []User
