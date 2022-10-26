package app

import (
	"log"

	"github.com/00Duck/wishr-api/database"
	"github.com/gorilla/mux"
)

type Env struct {
	db     *database.DB
	Router *mux.Router
	Log    *log.Logger
}

func New() *Env {
	env := &Env{
		Router: mux.NewRouter(),
		db:     &database.DB{},
		Log:    log.Default(),
	}
	env.db.Connect()
	env.routes()
	return env
}
