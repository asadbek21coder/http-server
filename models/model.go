package models

type User struct {
	ID    int
	Name  string
	Email string
}

type DeleteUserRequest struct {
	ID int
}

// DB
var Users []User
