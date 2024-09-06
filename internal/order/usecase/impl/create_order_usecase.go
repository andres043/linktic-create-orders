package impl

import (
	"context"
	"linktic-create-orders/internal/order/model"
	"linktic-create-orders/internal/order/port"
)

type CreateOrderUseCase struct {
	createOrderPort port.CreateOrderPort
}

func NewCreateOrder(createOrderPort port.CreateOrderPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		createOrderPort: createOrderPort,
	}
}

func (c *CreateOrderUseCase) Execute(ctx context.Context, order model.Order) error {
	return c.createOrderPort.CreateOrder(ctx, order)
}
