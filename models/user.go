package models

import "time"

type User struct {
	Name string
	Email string
	RegistrationDate time.Time
}
