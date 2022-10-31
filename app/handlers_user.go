package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

func (env *Env) HandleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		if ok := env.decodeRequest(w, r, &user); !ok {
			return
		}
		id, err := env.db.UserCreate(&user)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: id})
	}
}

func (env *Env) HandleUserRetrieveAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := env.db.UserRetrieveAll()
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: users})
	}
}

func (env *Env) HandleUserRetrieveOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user_id := params["id"]
		user, err := env.db.UserRetrieveOne(user_id)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: user})
	}
}

func (env *Env) HandleUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		if ok := env.decodeRequest(w, r, &user); !ok {
			return
		}
		err := env.db.UserUpdate(&user)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: user.ID})
	}
}

func (env *Env) HandleUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["id"]
		msg, err := env.db.UserDelete(userId)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: msg})
	}
}

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
