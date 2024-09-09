package application

import (
	"github.com/gorilla/mux"
	"linktic-create-orders/internal/order/adapter/input/web"
)

type Application struct {
	Router      *mux.Router
	CreateOrder web.CreateOrderHandler
}

func NewApplication(createOrder *web.CreateOrderHandler) *Application {
	app := &Application{
		Router: mux.NewRouter(),
	}

	app.Router.HandleFunc("/v1/orders", createOrder.Handle).Methods("POST")

	return app
}
