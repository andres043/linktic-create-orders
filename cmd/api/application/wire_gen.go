// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/google/wire"
	"linktic-create-orders/internal/order/adapter/input/web"
	"linktic-create-orders/internal/order/adapter/output"
	"linktic-create-orders/internal/order/port"
	"linktic-create-orders/internal/order/usecase"
	"linktic-create-orders/internal/order/usecase/impl"
)

// Injectors from wire.go:

func InitializeApplication() (*Application, error) {
	dynamoDB := output.NewDynamoDb()
	createOrderUseCase := impl.NewCreateOrder(dynamoDB)
	createOrderHandler := web.NewCreateOrderHandler(createOrderUseCase)
	application := NewApplication(createOrderHandler)
	return application, nil
}

// wire.go:

var orderInputAdapterSet = wire.NewSet(web.NewCreateOrderHandler)

var orderOutputAdapterSet = wire.NewSet(wire.Bind(new(port.CreateOrderPort), new(*output.DynamoDB)), output.NewDynamoDb)

var orderUseCaseSet = wire.NewSet(wire.Bind(new(usecase.CreateOrder), new(*impl.CreateOrderUseCase)), impl.NewCreateOrder)
