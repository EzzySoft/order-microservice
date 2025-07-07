package order_interface

import (
	"context"
	"order-service/internal/order/domain/entity"
)

type OrderRepository interface {
	Save(ctx context.Context, order entity.Order) error
	FindByID(ctx context.Context, id string) (entity.Order, error)
	AllOrderUIDs(ctx context.Context) ([]string, error)
}
