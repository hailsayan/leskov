package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/hailsayan/woland/auth"
	"github.com/hailsayan/woland/types"
	"github.com/hailsayan/woland/utils"
)

func (h *Server) ProductRegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.handleGetProduct).Methods(http.MethodGet)

	// admin routes
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.store.Users)).Methods(http.MethodPost)
}

func (h *Server) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	product, err := h.store.Product.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Server) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.store.Product.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Server) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.CreateProductPayload
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate().Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.Product.CreateProduct(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, product)
}
