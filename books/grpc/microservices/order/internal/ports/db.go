package ports

import "github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/domain"
type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}