package main

import (
	"net/http"
)

func AuthenticateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}
