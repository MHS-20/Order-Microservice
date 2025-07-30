package ports

import "github.com/MHS-20/Order-Microservice/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
