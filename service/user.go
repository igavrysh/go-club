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

func AddUser(name string, email string) (models.User, map[string]error) {
	var errs map[string]error

	if UserExists(email) {
		errs["general"] = errors.New("user with specified email already exists")
		return models.User{}, errs
	}

	var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9\\s.]*$")
	if !usernameRegex.MatchString(name) {
		errs["name"] = errors.New("user name must only have english letters, dots or spaces")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		errs["email"] = errors.New("user name must only have english letters, dots or spaces")
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

/*
type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

 */