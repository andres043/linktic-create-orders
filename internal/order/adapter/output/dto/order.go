package dto

import "linktic-create-orders/internal/order/model"

type Order struct {
	OrderID  string  `dynamodbav:"order_id"`
	Product  string  `dynamodbav:"product"`
	Quantity int     `dynamodbav:"quantity"`
	Price    float64 `dynamodbav:"price"`
}

func ToDto(order model.Order) Order {
	return Order{
		OrderID:  order.OrderID,
		Product:  order.Product,
		Quantity: order.Quantity,
		Price:    order.Price,
	}
}
