package app

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

// HandleLoginUser authenticates user
func (env *Env) HandleLoginUser() http.HandlerFunc {
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
		secureCookie := true
		if strings.ToUpper(os.Getenv("USE_SECURE_COOKIE")) != "TRUE" {
			secureCookie = false
		}
		sessionCookie := &http.Cookie{
			Name:   auth.SessionCookieName,
			Value:  session.ID,
			MaxAge: 60 * 60 * 24 * 30, // days
			// Expires:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			Secure:   secureCookie,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}

		profile := &models.ProfileUser{
			ID:       session.UserID,
			UserName: session.UserName,
			FullName: session.FullName,
		}

		http.SetCookie(w, sessionCookie)
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: profile})
	}
}

// LogOutUser send a trash cookie over token to force the user out of the system
func (env *Env) HandleLogOutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		SendLogoutCookie(w)
		env.encodeResponse(w, &ResponseModel{Message: "You have been logged out."})
	}
}

// HandleRegisterUser handles self service user creation
func (env *Env) HandleRegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		if ok := env.decodeRequest(w, r, &user); !ok {
			return
		}
		regEnabled := strings.ToUpper(os.Getenv("REGISTRATION_ENABLED"))
		if regEnabled != "TRUE" {
			env.encodeResponse(w, &ResponseModel{Message: "Registration is currently disabled."})
			return
		}
		regCode := os.Getenv("REGISTRATION_CODE")
		if regCode == "" && regEnabled == "TRUE" {
			env.Log.Println("Please set environment variable REGISTRATION_CODE to allow registration of intended users.")
			env.encodeResponse(w, &ResponseModel{Message: "Registration is currently disabled."})
			return
		}
		if regCode != user.RegistrationCode {
			env.encodeResponse(w, &ResponseModel{Message: "Registration code is invalid."})
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

// ValidateSessionMiddleware checks the token in the session cookie to ensure user is authenticated
func (env *Env) ValidateSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie(auth.SessionCookieName)
		if err != nil {
			env.Log.Println("Error retrieving cookie for validation: " + err.Error())
			errResponse(w, env.Log, http.StatusUnauthorized, errors.New("No cookie found"))
			return
		}
		session, err := env.db.CheckSession(sessionCookie.Value)
		if err != nil {
			env.Log.Println("Error on CheckSession: " + err.Error())
			SendLogoutCookie(w)
			errResponse(w, env.Log, http.StatusUnauthorized, errors.New("You are not logged in"))
			return
		}
		//add the session information to the request context
		r = r.WithContext(auth.NewContext(r.Context(), session))

		next.ServeHTTP(w, r)
	})
}

// ValidationCheck protected endpoint with middlware. Returns 200, but middleware will catch if there's a problem.
func (env *Env) ValidationCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env.encodeResponse(w, &ResponseModel{})
	}
}

// SendLogoutCookie separated out since used in multiple places
func SendLogoutCookie(w http.ResponseWriter) {
	logoutCookie := &http.Cookie{
		Name:     auth.SessionCookieName,
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
	}
	http.SetCookie(w, logoutCookie)
}

func (env *Env) HandleResetTokenValidationAndReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		token := params["token"]
		pw := ""
		if token == "" {
			errResponse(w, env.Log, http.StatusBadRequest, errors.New("Missing reset password token"))
			return
		}
		if r.Method == "POST" {
			pwModel := models.PasswordResetModel{}
			if ok := env.decodeRequest(w, r, &pwModel); !ok {
				return
			}
			pw = pwModel.Password
		}

		err := env.db.ValidateAndResetUser(token, pw)
		if err != nil {
			errResponse(w, env.Log, http.StatusUnauthorized, err)
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success"})
	}
}
