package app

import (
	"net/http"
)

func (env *Env) routes() {
	unprotected := env.Router.PathPrefix("/api/open").Subrouter()
	unprotected.HandleFunc("/login", func(http.ResponseWriter, *http.Request) { print("hello world") }).Methods("POST")
}
