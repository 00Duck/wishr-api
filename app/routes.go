package app

func (env *Env) routes() {
	unprotected := env.Router.PathPrefix("/api/open").Subrouter()
	unprotected.HandleFunc("/login", env.AuthenticateUser()).Methods("POST")
}
