package app

func (env *Env) routes() {
	unprotected := env.Router.PathPrefix("/api/open").Subrouter()
	unprotected.HandleFunc("/login", env.LoginUser()).Methods("POST")
	unprotected.HandleFunc("/logout", env.LogOutUser()).Methods("POST")
	unprotected.HandleFunc("/register", env.RegisterUser()).Methods("POST")

	protected := env.Router.PathPrefix("/api/prot").Subrouter()
	protected.Use(env.ValidateSessionMiddleware)
	protected.HandleFunc("/validate", env.ValidationCheck()).Methods("GET")
	protected.HandleFunc("/user", env.HandleUserRetrieveAll()).Methods("GET")
	protected.HandleFunc("/user/{id}", env.HandleUserRetrieveOne()).Methods("GET")
	protected.HandleFunc("/user", env.HandleUserCreate()).Methods("POST")
	protected.HandleFunc("/user", env.HandleUserUpdate()).Methods("PATCH")
	protected.HandleFunc("/user/{id}", env.HandleUserDelete()).Methods("DELETE")

	protected.HandleFunc("/wishlist", env.HandleWishlistRetrieveAll()).Methods("GET")
	protected.HandleFunc("/wishlist/{id}", env.HandleWishlistRetrieveOne()).Methods("GET")
	protected.HandleFunc("/wishlist", env.HandleWishlistCreate()).Methods("POST")
	protected.HandleFunc("/wishlist", env.HandleWishlistUpdate()).Methods("PATCH")
	protected.HandleFunc("/wishlist/{id}", env.HandleWishlistDelete()).Methods("DELETE")

}
