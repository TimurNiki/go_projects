package persistance

import "github.com/TimurNiki/go_api_tutorial/v6/src/application/domain/entities"

type orderRepository struct {
	db *gorm.DB
}

type OrderRepository interface {
	GetOrderById(id int64) *entities.Order
	CreateOrder(order entities.Order) *entities.Order
	ShipOrderByCargoCode(cargoCode string) error
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (or orderRepository) GetOrderById(order *entities.Order) *entities.Order {
	var order entities.Order
	result := or.db.Preload("OrderLineItems").First(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(result.Error)
	}

	return &order
}

func (or orderRepository) CreateOrder(order *entities.Order) *entities.Order {
	result := or.db.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	return &order
}

func (or orderRepository) ShipOrderByCargoCode(cargoCode string) error {
	result := or.db.Model(&entity.Order{}).Scopes(GetCargoCodeById(cargoCode), NonCancelledOrders).Update("is_shipped", true)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetCargoCodeByID(code string) func(db *gorm.DB) *gorm.DB {

}

func NonCancelledOrders(db *gorm.DB) *gorm.DB {
	return db.Where("is_cancelled = ?", false)
}
