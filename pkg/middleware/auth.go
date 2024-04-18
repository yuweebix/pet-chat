package middleware

import (
	"errors"
	"net/http"

	"github.com/yuweebix/pet-chat/pkg/repository"
	"gorm.io/gorm"
)

func IsAuthed(db *gorm.DB) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session_token")
			if err != nil {
				if errors.Is(err, http.ErrNoCookie) {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			clientToken := c.Value
			session, err := repository.GetSessionByToken(db, clientToken)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if session.IsAboutToExpire() {
				if err := repository.UpdateSessionExpiry(db, session); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Set the new session token in the cookie
				http.SetCookie(w, &http.Cookie{
					Name:     "session_token",
					Value:    session.Token,
					Expires:  session.ExpiresAt,
					Path:     "/",
					Secure:   false, // get rid of on htttps
					HttpOnly: false, // for javascript
				})
			}

			next.ServeHTTP(w, r)
		})
	}
}

func IsUnauthed(db *gorm.DB) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := repository.GetSessionByCookie(r, db)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// this functions seems to be useless, cuz I ain't receiving them cookies if they are expired :9
			if session.IsExpired() {
				if err := repository.DeleteSession(db, session); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "Already authenticated", http.StatusForbidden)
		})
	}
}
