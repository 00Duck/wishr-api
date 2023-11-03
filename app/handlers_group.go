package app

import (
	"net/http"

	"github.com/00Duck/wishr-api/models"
	"github.com/gorilla/mux"
)

func (env *Env) HandleGroupCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := models.Group{}
		if ok := env.decodeRequest(w, r, &group); !ok {
			return
		}
		id, err := env.db.GroupCreate(&group)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: id})
	}
}

// Not currently in use
// func (env *Env) HandleGroupRetrieveOne() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		group_id := params["id"]
// 		group, err := env.db.GroupRetrieveOne(group_id)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: group})
// 	}
// }

// Not currently in use
func (env *Env) HandleGroupUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := models.Group{}
		if ok := env.decodeRequest(w, r, &group); !ok {
			return
		}
		err := env.db.GroupUpdate(&group)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error()})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: group.ID})
	}
}

// Not currently in use
func (env *Env) HandleGroupDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		groupId := params["id"]
		msg, err := env.db.GroupDelete(groupId)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: msg})
	}
}

// Returns the groups for a user
func (env *Env) HandleGetGroupsForUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID := params["user"]
		groups, err := env.db.GetGroupsForUser(userID)
		if err != nil {
			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: groups})
			return
		}
		env.encodeResponse(w, &ResponseModel{Message: "success", Data: groups})
	}
}

// Returns a list of all selectable groups for the given wishlist (all groups minus already shared minus the requestor)
// func (env *Env) HandleGetSelectableShareGroupList() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		wishlistID := params["wishlist"]
// 		session := auth.FromContext(r.Context())
// 		groups, err := env.db.GetSelectableGroupsToShareList(session, wishlistID)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: groups})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: groups})
// 	}
// }

// // Sets the list of groups that a wishlist gets shared to
// func (env *Env) HandleSetSharedGroupsForWishlist() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		wishlistID := params["wishlist"]
// 		groups := []models.Group{}
// 		if ok := env.decodeRequest(w, r, &groups); !ok {
// 			return
// 		}
// 		err := env.db.SetGroupsForWishlist(wishlistID, groups)
// 		if err != nil {
// 			env.encodeResponse(w, &ResponseModel{Message: "Error: " + err.Error(), Data: nil})
// 			return
// 		}
// 		env.encodeResponse(w, &ResponseModel{Message: "success", Data: nil})
// 	}
// }
