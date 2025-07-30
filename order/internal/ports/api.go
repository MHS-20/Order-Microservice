package ports

import (
	"github.com/MHS-20/Order-Microservice/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
