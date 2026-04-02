package api

import (
	"database/sql"
	"fmt"
	"net/http"

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
	subrouter := router.PathPrefix("/api/v1").subrouter()

	return http.ListenAndServe(s.addr, router)
}

func (s *APIServer) InitializeDB() error {
	err := s.db.Ping()
	if err != nil {
		fmt.Errorf("error connecting to db....")
	}
	return nil
}
