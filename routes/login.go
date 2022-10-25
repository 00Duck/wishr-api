package routes

import (
	"net/http"

	"github.com/00Duck/wishr-api/wishr"
)

func (env *wishr.Env) AuthenticateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}
