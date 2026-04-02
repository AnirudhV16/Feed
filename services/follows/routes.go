package follows

import (
	"net/http"

	"github.com/AnirudhV16/Feed/services/auth"
	"github.com/AnirudhV16/Feed/types"
	"github.com/AnirudhV16/Feed/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserStore   types.UserStore
	FollowStore types.FolowStore
}

func NewHandler(userstore types.UserStore, followstore types.FolowStore) *Handler {
	return &Handler{UserStore: userstore, FollowStore: followstore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/follow", auth.WithJWTAuth(h.FollowHandler, h.UserStore)).Methods("POST")
}

func (h *Handler) FollowHandler(w http.ResponseWriter, r *http.Request) {
	//user gives username of the user he wants to follow
	//take that username and get the userid
	//add that is to the follows table

	//1.parse the request
	payload := new(types.FollowPayload)
	err := utils.JSONParse(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//get the followers id from the context
	followerid := auth.GetUserIDFromContext(r.Context())
	payload.FollowerId = followerid

	//store this relation in follows db
	err = h.FollowStore.AddFollower(*payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
