package ports

import ( 

	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/application/core/domain"
"context"
)
type APIPort interface {
	// PlaceOrder(order domain.Order) (domain.Order, error)
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
