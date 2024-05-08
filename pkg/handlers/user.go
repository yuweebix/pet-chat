package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/yuweebix/pet-chat/pkg/middleware"
	"github.com/yuweebix/pet-chat/pkg/models"
	"github.com/yuweebix/pet-chat/pkg/repository"
	"github.com/yuweebix/pet-chat/pkg/utils"
	"gorm.io/gorm"
)

var templates *template.Template
var db *gorm.DB

func initUserRoutes(mux *http.ServeMux, new_db *gorm.DB) error {
	var err error
	templates, err = utils.ParseTemplates("web/templates/user")
	if err != nil {
		return err
	}

	db = new_db

	// Routes that require the user to be unauthenticated
	unauthedMux := http.NewServeMux()
	unauthedMux.HandleFunc("GET /register/", registerGet)
	unauthedMux.HandleFunc("POST /register/", registerPost)
	unauthedMux.HandleFunc("GET /login/", loginGet)
	unauthedMux.HandleFunc("POST /login/", loginPost)
	mux.Handle("/unauthed/", http.StripPrefix("/unauthed", middleware.IsUnauthed(db)(unauthedMux)))

	// Routes that require the user to be authenticated
	authedMux := http.NewServeMux()
	authedMux.HandleFunc("GET /me/", me)
	authedMux.HandleFunc("PUT /me/update/", update)
	authedMux.HandleFunc("POST /me/logout/", logout)
	authedMux.HandleFunc("DELETE /me/delete/", delete)
	mux.Handle("/authed/", http.StripPrefix("/authed", middleware.IsAuthed(db)(authedMux)))

	// Public route
	mux.HandleFunc("/{username}/", profile)

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
	http.Redirect(w, r, "/home/", http.StatusSeeOther)
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

	session, err := repository.GetSessionByUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		Path:     "/",   // make the cookie available on all paths
		Secure:   false, // get rid of on htttps
		HttpOnly: false, // for javascript
	})

	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}

func profile(w http.ResponseWriter, r *http.Request) {
	// Implement your handler logic here
}

func me(w http.ResponseWriter, r *http.Request) {
	session, err := repository.GetSessionByCookie(r, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user := session.User
	data := map[string]interface{}{
		"Email":    user.Email,
		"Username": user.Username,
	}

	err = templates.ExecuteTemplate(w, "me.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedUser := models.UserCreate{
		Email:          r.FormValue("email"),
		Username:       r.FormValue("username"),
		Password:       r.FormValue("password"),
		RepeatPassword: r.FormValue("repeat_password"),
	}

	if !utils.CheckEmail(updatedUser.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if !utils.CheckUsername(updatedUser.Username) {
		http.Error(w, "Invalid username format", http.StatusBadRequest)
		return
	}
	if !utils.CheckPassword(updatedUser.Password) {
		http.Error(w, "Invalid password format", http.StatusBadRequest)
		return
	}
	if !utils.CheckPassword(updatedUser.RepeatPassword) {
		http.Error(w, "Invalid repeat_password format", http.StatusBadRequest)
		return
	}
	if updatedUser.Password != updatedUser.RepeatPassword {
		http.Error(w, "Passwords don't match", http.StatusBadRequest)
		return
	}

	session, err := repository.GetSessionByCookie(r, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := repository.UpdateUser(db, &updatedUser, &session.User); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := repository.GetSessionByCookie(r, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := repository.DeleteSession(db, session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the cookie's max age to -1 to delete it
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

func delete(w http.ResponseWriter, r *http.Request) {
	session, err := repository.GetSessionByCookie(r, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = repository.DeleteUser(db, session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logout(w, r)
}
