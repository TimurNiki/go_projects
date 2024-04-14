package types

import "context"

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context context.Context) []*orders.Order
}
