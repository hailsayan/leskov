package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hailsayan/woland/types"
	"github.com/hailsayan/woland/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get json payload
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if the user exists

	// if it doesn't exist, we create the user
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
