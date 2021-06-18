package controllers

import (
	"github.com/igavrysh/go-club/models"
	"github.com/igavrysh/go-club/service"
	"html/template"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	// process form submission
	var e map[string]error
	var un string
	var ue string
	var pendingUser = models.UserView{}
	if req.Method == http.MethodPost {
		// get form values
		un = req.FormValue("name")
		ue = req.FormValue("email")
		_, e = service.AddUser(un, ue)
		if e != nil {
			pendingUser = models.UserView{
				Number:           0,
				Name:             un,
				Email:            ue,
				RegistrationDate: "",
			}
		}

	}
	users := service.Users()
	var uvs = models.UsersView{}

	var uv []models.UserView
	for i, u := range users {
		uv = append(uv, models.UserView{
			Number:           i+1,
			Name:             u.Name,
			Email:            u.Email,
			RegistrationDate: u.RegistrationDate.Format("2006-01-02"),
		})
	}
	uvs = models.UsersView{
		PendingUser: pendingUser,
		Errors: e,
		Users: uv,
	}
	c.tpl.ExecuteTemplate(w, "index.gohtml", uvs)
}

func (c Controller) Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		service.ClearUsers()
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}
