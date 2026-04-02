package posts

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/AnirudhV16/Feed/services/auth"
	"github.com/AnirudhV16/Feed/services/users"
	"github.com/AnirudhV16/Feed/types"
	"github.com/AnirudhV16/Feed/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	PostStore Store
	UserStore users.Store
}

func NewHandler(PostStore Store, UserStore users.Store) Handler {
	return Handler{PostStore: PostStore,
		UserStore: UserStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/post", auth.WithJWTAuth(h.UploadPostHandler, h.UserStore)).Methods("POST")
	router.HandleFunc("/feed", h.FeedHandler).Methods("GET")
}

func (h *Handler) UploadPostHandler(w http.ResponseWriter, r *http.Request) {

	//get the post details from user request (img or text)
	file, meta, err := r.FormFile("image")
	text := r.FormValue("text")

	var imgUrl string
	if err == nil {
		file_path := "./uploads/" + meta.Filename

		emp_file, err := os.Create(file_path)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(emp_file, file)
		imgUrl = "/uploads" + meta.Filename
	}

	if text == "" && imgUrl == "" {
		utils.WriteError(w, 400, fmt.Errorf("empty post"))
		return
	}

	//store imgurl in the database
	//so need a post store first
	//need a post type first
	post := new(types.Post{
		Content: text,
		ImgUrl:  imgUrl,
	})
	err = h.PostStore.CreatePost(r, *post)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) FeedHandler(w http.ResponseWriter, r *http.Request) {

	//got the slice of posts from the postStore
	//return the post details as response
	id := auth.GetUserIDFromContext(r.Context())
	posts, err := h.PostStore.GetPosts(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, 200, posts)
}
