package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	userMux := http.NewServeMux()
	if err := InitUserRoutes(userMux, db); err != nil {
		return nil, err
	}
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

	return mux, nil
}
