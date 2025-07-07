package application

import (
	"context"
	"log"
	"order-service/internal/order/domain/factory"
	iface "order-service/internal/order/interface"
)

type OrderService struct {
	Repo iface.OrderRepository
}

func (s *OrderService) HandleOrderJSON(data []byte) error {
	order, err := factory.OrderFromJSON(data)
	if err != nil {
		log.Printf("parse order err: %v", err)
		return err
	}
	log.Printf("[parser] ✅ Order parsed [%+v]", order.UID())
	err = s.Repo.Save(context.Background(), order)
	if err != nil {
		log.Printf("save order err: %v", err)
		return err
	}
	log.Printf("[parser] ✅ Order saved ")
	return nil
}
