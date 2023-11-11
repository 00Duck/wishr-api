package app

import (
	"log"

	"github.com/00Duck/wishr-api/cmd"
	"github.com/00Duck/wishr-api/database"
	"github.com/gorilla/mux"
)

type Env struct {
	db     *database.DB
	Router *mux.Router
	Log    *log.Logger
	CLI    *cmd.CLI
}

func New() *Env {
	env := &Env{
		Router: mux.NewRouter(),
		db:     &database.DB{},
		Log:    log.Default(),
		CLI:    nil,
	}

	env.db.Connect()
	env.routes()
	env.CLI = cmd.NewCLI(env.db)
	return env
}
