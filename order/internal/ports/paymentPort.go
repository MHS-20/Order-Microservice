package ports

import "github.com/MHS-20/Order-Microservice/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
