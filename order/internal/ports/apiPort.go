package ports

import (
	"context"

	"github.com/MHS-20/Order-Microservice/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
