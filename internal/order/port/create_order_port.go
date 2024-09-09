package port

import (
	"context"
	"linktic-create-orders/internal/order/model"
)

type CreateOrderPort interface {
	CreateOrder(ctx context.Context, order model.Order) error
}
