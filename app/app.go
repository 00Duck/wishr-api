package app

import (
	"github.com/00Duck/wishr-api/database"
	"github.com/gorilla/mux"
)

type Env struct {
	DB     *database.DB
	Router *mux.Router
}

func New() *Env {
	env := &Env{
		Router: mux.NewRouter(),
	}
	env.routes()
	return env
}
