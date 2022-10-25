package wishr

import (
	"net/http"

	wishr "github.com/00Duck/wishr-api"
)

func (env *wishr.Env) authenticateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}
