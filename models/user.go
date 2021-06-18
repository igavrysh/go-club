package models

import "time"

type User struct {
	Name string
	Email string
	RegistrationDate time.Time
}

type UserView struct {
	Number int
	Name string
	Email string
	RegistrationDate string
}

type UsersView struct {
	PendingUser UserView
	Errors map[string]error
	Users []UserView
}
