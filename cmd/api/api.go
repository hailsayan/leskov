package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hailsayan/woland/internal/store"
)

type Server struct {
	addr  string
	db    *sql.DB
	store store.Storage
}

func NewServer (addr string, db *sql.DB, store store.Storage) *Server{
	return &Server{
		addr: addr,
		db: db,
		store: store,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	s.UserRegisterRoutes(subrouter)
	s.ProductRegisterRoutes(subrouter)
	s.CartRegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
