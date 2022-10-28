package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/models"
)

func (env *Env) AuthenticateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := models.AuthenticateModel{}
		if ok := env.decodeRequest(w, r, &auth); !ok {
			return
		}
		test := env.db.Authenticate(&auth)
		env.encodeResponse(w, &ResponseModel{Message: test})
	}
}
