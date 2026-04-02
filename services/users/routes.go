package users

import (
	"fmt"
	"net/http"

	"github.com/AnirudhV16/Feed/services/auth"
	"github.com/AnirudhV16/Feed/types"
	"github.com/AnirudhV16/Feed/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	userStore *Store
}

func NewHandler(s *Store) *Handler {
	return &Handler{userStore: s}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", h.LoginHandler).Methods("POST")
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//parsing register request body
	payload := new(types.RegisterPayload)
	err := utils.JSONParse(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("badrequest...."))
		return
	}

	//check if user exists in db
	_, err = h.userStore.GetUserByGmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("duplicate !! user already exists...."))
		return
	}

	//hash password and store the name and hash in db
	//first need db
	hash, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Internal server errorrr....."))
		return
	}
	//got the hashed password store it in db
	//need the userstore dependency
	user := new(types.User{
		FirstName: payload.FirstName,
		Email:     payload.Email,
		Password:  hash,
	})
	err = h.userStore.CreateUser(user)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, nil)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//parsing register request body
	payload := new(types.LoginPayload)
	err := utils.JSONParse(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("badrequest...."))
		return
	}

	//get user details from db
	u, err := h.userStore.GetUserByGmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no records found"))
		return
	}

	status := auth.Compare(u.Password, []byte(payload.Password))
	if status == false {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad credentials try again"))
		return
	}

}
