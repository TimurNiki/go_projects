package ports

import "github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/payment/internal/application/core/domain"
type PaymentPort interface {
	Charge(*domain.Order) error
   }