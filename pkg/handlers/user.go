package handlers

import (
	"html/template"
	"net/http"

	"github.com/yuweebix/pet-chat/pkg/models"
	"github.com/yuweebix/pet-chat/pkg/utils"
)

var templates *template.Template

func InitUserRoutes(mux *http.ServeMux) error {
	var err error
	templates, err = utils.ParseTemplates("web/templates/user")
	if err != nil {
		return err
	}
	mux.HandleFunc("GET /register/", registerGet)
	mux.HandleFunc("POST /register/", registerPost)
	mux.HandleFunc("POST /login/", login)
	mux.HandleFunc("POST /logout", logout)
	mux.HandleFunc("GET /get/", get)
	mux.HandleFunc("PUT /update/", update)
	mux.HandleFunc("DELETE /delete/", delete)

	return nil
}

func registerGet(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func registerPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdUser := models.UserCreate{
		Email:    r.FormValue(("email")),
		Username: r.FormValue(("username")),
		Password: r.FormValue(("password")),
	}

	if !utils.CheckEmail(createdUser.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if !utils.CheckUsername(createdUser.Username) {
		http.Error(w, "Invalid username format", http.StatusBadRequest)
		return
	}
	if !utils.CheckPassword(createdUser.Password) {
		http.Error(w, "Invalid password format", http.StatusBadRequest)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func logout(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func get(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func update(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func delete(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}
