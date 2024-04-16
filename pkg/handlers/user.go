package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/yuweebix/pet-chat/pkg/models"
	"github.com/yuweebix/pet-chat/pkg/repository"
	"github.com/yuweebix/pet-chat/pkg/utils"
	"gorm.io/gorm"
)

var templates *template.Template
var db *gorm.DB

func InitUserRoutes(mux *http.ServeMux, new_db *gorm.DB) error {
	var err error
	templates, err = utils.ParseTemplates("web/templates/user")
	if err != nil {
		return err
	}

	db = new_db

	mux.HandleFunc("GET /register/", registerGet)
	mux.HandleFunc("POST /register/", registerPost)
	mux.HandleFunc("GET /login/", loginGet)
	mux.HandleFunc("POST /login/", loginPost)

	mux.HandleFunc("GET /{username}/", profile)

	mux.HandleFunc("GET /me/", me)
	mux.HandleFunc("PUT /update/", update)
	mux.HandleFunc("POST /logout/", logout)
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
		Email:          r.FormValue("email"),
		Username:       r.FormValue("username"),
		Password:       r.FormValue("password"),
		RepeatPassword: r.FormValue("repeat_password"),
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
	if !utils.CheckPassword(createdUser.RepeatPassword) {
		http.Error(w, "Invalid repeat_password format", http.StatusBadRequest)
		return
	}
	if createdUser.Password != createdUser.RepeatPassword {
		http.Error(w, "Passwords don't match", http.StatusBadRequest)
		return
	}

	if err := repository.CreateUser(db, &createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	loggedUser := models.UserLogin{
		UsernameOrEmail: r.FormValue("username_or_email"),
		Password:        r.FormValue("password"),
	}

	if !utils.CheckEmail(loggedUser.UsernameOrEmail) && !utils.CheckUsername(loggedUser.UsernameOrEmail) {
		http.Error(w, "Invalid username_or_email format", http.StatusBadRequest)
		return
	}

	user, err := repository.LoginUser(db, &loggedUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else if errors.Is(err, repository.ErrInvalidPassword) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := repository.CreateSession(db, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session, err := repository.GetSession(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   session.Token,
		Expires: session.ExpiresAt,
	})
}

func profile(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func me(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func update(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func logout(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func delete(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}
