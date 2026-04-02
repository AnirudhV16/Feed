package api

import (
	"database/sql"
	"net/http"

	"github.com/AnirudhV16/Feed/services/follows"
	"github.com/AnirudhV16/Feed/services/posts"
	"github.com/AnirudhV16/Feed/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//user handler
	UserStore := users.NewStore(s.db)
	UserHandler := users.NewHandler(UserStore)
	UserHandler.RegisterRoutes(subrouter)

	//follow handler
	FollowStore := follows.NewStore(s.db)
	FollowHandler := follows.NewHandler(UserStore, FollowStore)
	FollowHandler.RegisterRoutes(subrouter)

	//feed handler
	PostStore := posts.NewStore(s.db)
	PostHandler := posts.NewHandler(PostStore, UserStore)
	PostHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
