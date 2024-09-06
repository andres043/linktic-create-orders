package web

import (
	"encoding/json"
	"linktic-create-orders/internal/order/adapter/input/web/dto"
	"linktic-create-orders/internal/order/usecase"
	"net/http"
)

type CreateOrderHandler struct {
	createOrderUseCase usecase.CreateOrder
}

func NewCreateOrderHandler(createOrderUseCase usecase.CreateOrder) *CreateOrderHandler {
	return &CreateOrderHandler{
		createOrderUseCase: createOrderUseCase,
	}
}

// Handle CreateOrderHandler godoc
// @Summary Create an order
// @Description Create an order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body Order true "Order"
// @Success 200 {object} Order
// @Router /orders [post]
func (c *CreateOrderHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orderDto := dto.Order{}

	err := json.NewDecoder(r.Body).Decode(&orderDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = c.createOrderUseCase.Execute(ctx, orderDto.ToModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
