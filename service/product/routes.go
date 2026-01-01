package product

import (
	"net/http"

	"github.com/kasariks/api_golang/types"
	"github.com/kasariks/api_golang/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /products", h.handleCreateProduct)
	router.HandleFunc("GET /products", h.handleCreateProduct)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
