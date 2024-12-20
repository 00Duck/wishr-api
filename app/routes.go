package app

func (env *Env) routes() {

	unprotected := env.Router.PathPrefix("/api/open").Subrouter()
	unprotected.HandleFunc("/login", env.HandleLoginUser()).Methods("POST")
	unprotected.HandleFunc("/logout", env.HandleLogOutUser()).Methods("POST")
	unprotected.HandleFunc("/register", env.HandleRegisterUser()).Methods("POST")

	unprotected.HandleFunc("/passwordreset/{token}", env.HandleResetTokenValidationAndReset()).Methods("GET", "POST")

	protected := env.Router.PathPrefix("/api/prot").Subrouter()
	protected.Use(env.ValidateSessionMiddleware)
	protected.HandleFunc("/validate", env.ValidationCheck()).Methods("GET")
	// protected.HandleFunc("/user", env.HandleUserRetrieveAll()).Methods("GET")
	// protected.HandleFunc("/user/{id}", env.HandleUserRetrieveOne()).Methods("GET")
	// protected.HandleFunc("/user", env.HandleUserCreate()).Methods("POST")
	// protected.HandleFunc("/user", env.HandleUserUpdate()).Methods("PATCH")
	// protected.HandleFunc("/user/{id}", env.HandleUserDelete()).Methods("DELETE")
	protected.HandleFunc("/user/shared/selectable/{wishlist}", env.HandleGetSelectableShareUserList()).Methods("GET")
	protected.HandleFunc("/user/shared/{wishlist}", env.HandleGetSharedUsersForWishlist()).Methods("GET")
	protected.HandleFunc("/user/shared/{wishlist}", env.HandleSetSharedUsersForWishlist()).Methods("POST")

	protected.HandleFunc("/profile", env.HandleRetrieveProfile()).Methods("GET")

	protected.HandleFunc("/wishlist/shared", env.HandleWishlistRetrieveShared()).Methods("GET")
	protected.HandleFunc("/wishlist", env.HandleWishlistRetrieveAll()).Methods("GET")
	protected.HandleFunc("/wishlist/browse", env.HandleWishlistBrowse()).Methods("GET")
	protected.HandleFunc("/wishlist/{id}", env.HandleWishlistRetrieveOne()).Methods("GET")
	protected.HandleFunc("/wishlist", env.HandleWishlistUpsert()).Methods("POST")
	protected.HandleFunc("/wishlist/{id}", env.HandleWishlistDelete()).Methods("DELETE")

	protected.HandleFunc("/wishlist_item/reserve", env.HandleWishlistItemReserve()).Methods("POST")
	protected.HandleFunc("/wishlist_item/unreserve", env.HandleWishlistItemUnreserve()).Methods("POST")

	unprotected.HandleFunc("/images/{table}/{id}", env.HandleRetrieveImage()).Methods("GET")
	protected.HandleFunc("/images/{table}/upload", env.HandleImageUpload()).Methods("POST")
	protected.HandleFunc("/images/delete", env.HandleDeleteImage()).Methods("POST")

	// env.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	tpl, _ := route.GetPathTemplate()
	// 	met, _ := route.GetMethods()
	// 	fmt.Println(tpl, met)
	// 	return nil
	// })

}
