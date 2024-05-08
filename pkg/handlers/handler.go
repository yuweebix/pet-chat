package handlers

import (
	"html/template"
	"net/http"

	"github.com/yuweebix/pet-chat/pkg/middleware"
	"github.com/yuweebix/pet-chat/pkg/utils"
	"gorm.io/gorm"
)

var indexTemplate *template.Template
var homeTemplate *template.Template

func NewRouter(db *gorm.DB) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	if err := initIndexMux(mux, db); err != nil {
		return nil, err
	}

	if err := initHomeMux(mux, db); err != nil {
		return nil, err
	}

	userMux := http.NewServeMux()
	if err := initUserRoutes(userMux, db); err != nil {
		return nil, err
	}
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

	return mux, nil
}

func initIndexMux(mux *http.ServeMux, new_db *gorm.DB) error {
	var err error
	indexTemplate, err = utils.ParseTemplates("web/templates/index")
	if err != nil {
		return err
	}

	db = new_db

	indexMux := http.NewServeMux()
	indexMux.HandleFunc("GET /", indexGet)

	mux.Handle("/", middleware.IsUnauthed(db)(indexMux))

	return nil
}

func initHomeMux(mux *http.ServeMux, new_db *gorm.DB) error {
	var err error
	homeTemplate, err = utils.ParseTemplates("web/templates/home")
	if err != nil {
		return err
	}

	db = new_db

	homeMux := http.NewServeMux()
	homeMux.HandleFunc("GET /home/", homeGet)

	mux.Handle("/home/", middleware.IsAuthed(db)(homeMux))

	return nil
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
