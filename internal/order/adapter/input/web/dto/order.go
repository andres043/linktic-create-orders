package dto

import "linktic-create-orders/internal/order/model"

type Order struct {
	OrderID  string  `json:"order_id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (o *Order) ToModel() model.Order {
	return model.Order{
		Product:  o.Product,
		Quantity: o.Quantity,
		Price:    o.Price,
	}
}
