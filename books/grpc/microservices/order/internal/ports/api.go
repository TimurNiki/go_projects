package ports

import "github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
