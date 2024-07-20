package card

import (
	"fmt"
	"net/http"

	"github.com/TimurNiki/go_api_tutorial/v4/services/auth"
	"github.com/TimurNiki/go_api_tutorial/v4/types"
	"github.com/TimurNiki/go_api_tutorial/v4/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store      types.ProductStore
	orderStore types.OrderStore
	userStore  types.UserStore
}

func NewHandler(store types.ProductStore, orderStore types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:      store,
		orderStore: orderStore,
		userStore:  userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	//
	var card types.CartCheckoutPayload

	// we need to parse the request body before validating it
	// validating - checking if the payload is valid
	if err := utils.ParseJSON(r, &card); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// explanation: we need to validate the payload before processing it
	if err := utils.Validate.Struct(card); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// get productIds
	productIds, err := getCartItemsIDs(card.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products
	products, err := h.store.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// create order
	orderID, totalPrice, err := h.createOrder(products, card.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})

}
