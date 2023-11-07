package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (env *Env) HandleRetrieveImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		table := params["table"]
		id := params["id"]
		url := "./images/" + table + "/" + id

		http.ServeFile(w, r, url)
	}
}
