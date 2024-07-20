package api

import (
	"context"

	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/payment/internal/application/core/domain"
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/payment/internal/application/core/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) Charge(context context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
