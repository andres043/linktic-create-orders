package usecase

import (
	"context"
	"linktic-create-orders/internal/order/model"
)

type CreateOrder interface {
	Execute(ctx context.Context, order model.Order) error
}
