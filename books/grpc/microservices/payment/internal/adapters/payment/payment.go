package payment

import (
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/microservices-proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/payment/internal/core/domain"
	"context"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceURL string) (*Adapter, error) {
	var opts []grpc.DialOption // Data model for connection configurations
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials())) // This is for disabling TLS
	conn, err := grpc.Dial(paymentServiceUrl, opts...) // Connects to service
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := payment.NewPaymentClient(conn) // Initializes the new payment stub instance

	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(),
		&payment.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice(),
		})
	return err
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}
