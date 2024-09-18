package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hailsayan/woland/internal/store"
	"go.uber.org/zap"
)

type Server struct {
	addr  string
	db    *sql.DB
	store store.Storage
	logger *zap.SugaredLogger
}

func NewServer (addr string, db *sql.DB, store store.Storage, logger *zap.SugaredLogger) *Server{
	return &Server{
		addr: addr,
		db: db,
		store: store,
		logger: logger,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	s.UserRegisterRoutes(subrouter)
	s.ProductRegisterRoutes(subrouter)
	s.CartRegisterRoutes(subrouter)

	s.logger.Infow("Listening on", "addr", s.addr)
	return http.ListenAndServe(s.addr, router)
}
