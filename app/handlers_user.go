package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/auth"
	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

// Not currently in use
// func (env *Env) HandleUserCreate() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user := models.User{}
// 		if ok := env.decodeRequest(w, r, &user); !ok {
// 			return
// 		}
// 		id, err := env.db.UserCreate(&user)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: id})
// 	}
// }

// Not currently in use
// func (env *Env) HandleUserRetrieveAll() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := env.db.UserRetrieveAll()
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: users})
// 	}
// }

// Not currently in use
// func (env *Env) HandleUserRetrieveOne() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		user_id := params["id"]
// 		user, err := env.db.UserRetrieveOne(user_id)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: user})
// 	}
// }

// Not currently in use
// func (env *Env) HandleUserUpdate() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user := models.User{}
// 		if ok := env.decodeRequest(w, r, &user); !ok {
// 			return
// 		}
// 		err := env.db.UserUpdate(&user)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: user.ID})
// 	}
// }

// Not currently in use
// func (env *Env) HandleUserDelete() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		userId := params["id"]
// 		msg, err := env.db.UserDelete(userId)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: msg})
// 	}
// }

// Returns the already shared users for a wishlist
func (env *Env) HandleGetSharedUsersForWishlist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlistID := params["wishlist"]
		users, err := env.db.GetUsersForWishlist(wishlistID)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: users})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: users})
	}
}

// Returns a list of all selectable users for the given wishlist (all users minus already shared minus the requestor)
func (env *Env) HandleGetSelectableShareUserList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlistID := params["wishlist"]
		session := auth.FromContext(r.Context())
		users, err := env.db.GetSelectableUsersToShareList(session, wishlistID)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: users})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: users})
	}
}

// Sets the list of users that a wishlist gets shared to
func (env *Env) HandleSetSharedUsersForWishlist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		wishlistID := params["wishlist"]
		users := []models.User{}
		if ok := env.decodeRequest(w, r, &users); !ok {
			return
		}
		err := env.db.SetUsersForWishlist(wishlistID, users)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: nil})
	}
}

// Gets the users profile
func (env *Env) HandleRetrieveProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := auth.FromContext(r.Context())
		user, err := env.db.RetrieveProfile(session)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: user})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: user})
	}
}
