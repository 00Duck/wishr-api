package app

import (
	"net/http"
	"time"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
)

func (env *Env) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := models.LoginModel{}
		if ok := env.decodeRequest(w, r, &login); !ok {
			return
		}
		session, err := env.db.Authenticate(&login)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: err.Error()})
			return
		}
		sessionCookie := &http.Cookie{
			Name:  auth.SessionCookieName,
			Value: session.ID,
			// Note that core_session has TTL turned on - the session data will automatically be deleted from the database
			// once it gets stale. MaxAge should match time to live to prevent unexpected logouts.
			MaxAge:   60 * 60 * 24 * 14, // days
			Secure:   false,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, sessionCookie)
		env.encodeResponse(w, &ResponseModel{Message: "success"})
	}
}

// LogOutUser send a trash cookie over token to force the user out of the system
func (env *Env) LogOutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logoutCookie := &http.Cookie{
			Name:     auth.SessionCookieName,
			Value:    "",
			MaxAge:   -1,
			Expires:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Secure:   false,
			HttpOnly: true,
			Path:     "/",
		}

		sessionCookie, err := r.Cookie(auth.SessionCookieName)
		if err != nil {
			env.Log.Println("LogOutUser error (user has no session cookie): " + err.Error())
			env.encodeResponse(w, &ResponseModel{Message: "It appears you are already logged out."})
			return
		}

		err = env.db.Deauthenticate(sessionCookie.Value)
		if err != nil {
			env.Log.Println("Could not deauthenticate user: " + err.Error())
			env.encodeResponse(w, &ResponseModel{Message: "There was a problem logging you out. Please try again."})
			return
		}

		http.SetCookie(w, logoutCookie)
		env.encodeResponse(w, &ResponseModel{Message: "You have been logged out."})
	}
}

func (env *Env) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		if ok := env.decodeRequest(w, r, &user); !ok {
			return
		}
		err := env.db.Register(&user)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success"})
	}
}
