package serializer

import (
	"order-service/internal/order/application/api/scheme"
	"order-service/internal/order/domain/entity"
	"time"
)

func OrderToResponse(order entity.Order) scheme.OrderResponse {
	// Delivery
	delivery := order.Delivery()
	dResp := scheme.DeliveryResponse{
		Name:    delivery.Name(),
		Phone:   string(delivery.Phone()),
		Zip:     delivery.Zip(),
		City:    delivery.City(),
		Address: delivery.Address(),
		Region:  delivery.Region(),
		Email:   string(delivery.Email()),
	}
	// Payment
	payment := order.Payment()
	pResp := scheme.PaymentResponse{
		Transaction:  payment.Transaction(),
		RequestID:    payment.RequestID(),
		Currency:     string(payment.Currency()),
		Provider:     string(payment.Provider()),
		Amount:       payment.Amount().Int(),
		PaymentDT:    payment.PaymentDT(),
		Bank:         payment.Bank(),
		DeliveryCost: payment.DeliveryCost().Int(),
		GoodsTotal:   payment.GoodsTotal().Int(),
		CustomFee:    payment.CustomFee().Int(),
	}
	// Items
	var itemsResp []scheme.ItemResponse
	for _, item := range order.Items() {
		itemsResp = append(itemsResp, scheme.ItemResponse{
			ChrtID:      item.ChrtID(),
			TrackNumber: string(item.TrackNumber()),
			Price:       item.Price().Int(),
			RID:         string(item.RID()),
			Name:        item.Name(),
			Sale:        item.Sale(),
			Size:        item.Size(),
			TotalPrice:  item.TotalPrice().Int(),
			NmID:        item.NmID(),
			Brand:       item.Brand(),
			Status:      item.Status(),
		})
	}
	return scheme.OrderResponse{
		OrderUID:          string(order.UID()),
		TrackNumber:       string(order.TrackNumber()),
		Entry:             string(order.Entry()),
		Delivery:          dResp,
		Payment:           pResp,
		Items:             itemsResp,
		Locale:            string(order.Locale()),
		InternalSignature: order.InternalSignature(),
		CustomerID:        order.CustomerID(),
		DeliveryService:   string(order.DeliveryService()),
		ShardKey:          order.ShardKey(),
		SmID:              order.SmID(),
		DateCreated:       order.DateCreated().Format(time.RFC3339),
		OofShard:          order.OofShard(),
	}
}
