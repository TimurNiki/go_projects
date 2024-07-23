package api

import (
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/application/core/domain"
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/ports"
	"context"
)

type Application struct {
	db ports.DBPort // The API depends on DBPORT

}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

// func (a Application) PlaceOrder2(order domain.Order) (domain.Order, error) {
// 	err := a.db.Save(ctx,order)
// 	if err != nil {
// 		return domain.Order{}, err
// 	}
// 	return order, nil
// }


func (a Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.db.Get(ctx, id)
}