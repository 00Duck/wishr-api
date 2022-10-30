package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

func (env *Env) HandleWishlistUpsert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wishlist := models.Wishlist{}
		if ok := env.decodeRequest(w, r, &wishlist); !ok {
			return
		}
		session := auth.FromContext(r.Context())
		if wishlist.Owner == "" {
			wishlist.Owner = session.UserID
		}
		id, err := env.db.WishlistUpsert(&wishlist)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: id})
	}
}

func (env *Env) HandleWishlistRetrieveAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := auth.FromContext(r.Context())
		wishlists, err := env.db.WishlistRetrieveAll(session)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: wishlists})
	}
}

func (env *Env) HandleWishlistRetrieveShared() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := auth.FromContext(r.Context())
		wishlists, err := env.db.GetSharedWishlists(session)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: wishlists})
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
