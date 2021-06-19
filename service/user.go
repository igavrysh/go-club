package service

import (
	"errors"
	"github.com/igavrysh/go-club/models"
	"net/mail"
	"regexp"
	"time"
)

var usersDB = map[string]models.User{}

func UserExists(email string) bool {
	_, exists := usersDB[email]
	return exists
}

func UsersLen() int {
	return len(usersDB)
}

func AddUser(name string, email string) (models.User, map[string]error) {
	var errs map[string]error = map[string]error{}

	if UserExists(email) {
		errs["general"] = errors.New("user with specified email already exists")
		return models.User{}, errs
	}

	var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9\\s.]+$")
	if !usernameRegex.MatchString(name) {
		errs["name"] = errors.New("user name must be non-empty and only have english letters, dots or spaces")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		errs["email"] = err
	}

	if errs != nil && len(errs) > 0 {
		return models.User{}, errs
	}

	u := models.User{Name: name, Email: email, RegistrationDate: time.Now() }
	usersDB[email] = u
	return u, nil
}

func ClearUsers() {
	for email := range usersDB {
		delete(usersDB, email)
	}
}

func Users() []models.User {
	var r []models.User
	for _, v := range usersDB {
		r = append(r, v)
	}
	return r
}
