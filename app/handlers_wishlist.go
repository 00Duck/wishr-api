package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

// Handles both the creation and edit/update of wishlists
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

// Retrieves all wishlists for the given session user
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

// Retrieves all shared wishlists for the given session user
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

// Retrieves a wishlist. Used for both shared and owned wishlists
func (env *Env) HandleWishlistRetrieveOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlist_id := params["id"]
		session := auth.FromContext(r.Context())
		wishlist, err := env.db.WishlistRetrieveOne(session, wishlist_id)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: wishlist})
	}
}

// Deletes a wishlist
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

// HandleWishlistItemReserve reserves a Wishlist Item
func (env *Env) HandleWishlistItemReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wlItem := models.WishlistItem{}
		if ok := env.decodeRequest(w, r, &wlItem); !ok {
			return
		}
		session := auth.FromContext(r.Context())
		err := env.db.ReserveWishlistItem(session, &wlItem)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: nil})
	}
}

// HandleWishlistItemReserve reserves a Wishlist Item
func (env *Env) HandleWishlistItemUnreserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wlItem := models.WishlistItem{}
		if ok := env.decodeRequest(w, r, &wlItem); !ok {
			return
		}
		session := auth.FromContext(r.Context())
		err := env.db.UnreserveWishlistItem(session, &wlItem)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: nil})
	}
}
