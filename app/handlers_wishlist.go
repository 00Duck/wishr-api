package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

func (env *Env) HandleWishlistCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wishlist := models.Wishlist{}
		if ok := env.decodeRequest(w, r, &wishlist); !ok {
			return
		}
		id, err := env.db.WishlistCreate(&wishlist)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: id})
	}
}

func (env *Env) HandleWishlistRetrieveAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wishlists, err := env.db.WishlistRetrieveAll()
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: wishlists})
	}
}

func (env *Env) HandleWishlistRetrieveOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlist_id := params["id"]
		wishlist, err := env.db.WishlistRetrieveOne(wishlist_id)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: wishlist})
	}
}

func (env *Env) HandleWishlistUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wishlist := models.Wishlist{}
		if ok := env.decodeRequest(w, r, &wishlist); !ok {
			return
		}
		err := env.db.WishlistUpdate(&wishlist)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: wishlist.ID})
	}
}

func (env *Env) HandleWishlistDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlist_id := params["id"]
		msg, err := env.db.WishlistDelete(wishlist_id)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: msg})
	}
}
