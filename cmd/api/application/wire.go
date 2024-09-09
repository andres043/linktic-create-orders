//go:build wireinject
// +build wireinject

package application

import (
	"github.com/google/wire"
	"linktic-create-orders/internal/order/adapter/input/web"
	"linktic-create-orders/internal/order/adapter/output"
	"linktic-create-orders/internal/order/port"
	"linktic-create-orders/internal/order/usecase"
	"linktic-create-orders/internal/order/usecase/impl"
)

var orderInputAdapterSet = wire.NewSet(
	web.NewCreateOrderHandler,
)

var orderOutputAdapterSet = wire.NewSet(
	wire.Bind(new(port.CreateOrderPort), new(*output.DynamoDB)),
	output.NewDynamoDb,
)

var orderUseCaseSet = wire.NewSet(
	wire.Bind(new(usecase.CreateOrder), new(*impl.CreateOrderUseCase)),
	impl.NewCreateOrder,
)

func InitializeApplication() (*Application, error) {
	wire.Build(
		NewApplication,
		orderInputAdapterSet,
		orderOutputAdapterSet,
		orderUseCaseSet,
	)

	return &Application{}, nil
}
