package product

import (
	"fmt"
	"github.com/TimurNiki/go_api_tutorial/v4/types"
	"github.com/TimurNiki/go_api_tutorial/v4/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.handleGetProduct).Methods(http.MethodGet)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product id"))
		return
	}
	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}
	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}
