package services

import "github.com/TimurNiki/go_api_tutorial/v6/src/application/domain/entities"

type orderservice struct {
	orderRepository persistance.OrderRepository
}

func NewOrderService(orderRepository persistance.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}
type OrderService interface {
	CreateOrder(command model.CreateOrderCommand) *entities.Order
	GetOrderById(id int64) *entities.Order
	ShipOrderByCargoCode(cargoCode string) error
}

func (s orderservice) CreateOrder(command model.CreateOrderCommand) *entities.Order {
	order := model.MapToOrder(command)
	return service.orderRepository.CreateOrder(order)
}

func (s orderservice) GetOrderById(id int64) *entities.Order {
	return s.orderRepository.GetOrderById(id)
}

func (s orderservice) ShipOrderByCargoCode(cargoCode string) error {
	return s.orderRepository.ShipOrderByCargoCode(cargoCode)
}