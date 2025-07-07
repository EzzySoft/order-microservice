package factory

import (
	"encoding/json"
	"order-service/internal/order/domain/entity"
	"order-service/internal/order/domain/valueobject"
	"time"
)

type orderDTO struct {
	OrderUID          string      `json:"order_uid"`
	TrackNumber       string      `json:"track_number"`
	Entry             string      `json:"entry"`
	Delivery          deliveryDTO `json:"delivery"`
	Payment           paymentDTO  `json:"payment"`
	Items             []itemDTO   `json:"items"`
	Locale            string      `json:"locale"`
	InternalSignature string      `json:"internal_signature"`
	CustomerID        string      `json:"customer_id"`
	DeliveryService   string      `json:"delivery_service"`
	ShardKey          string      `json:"shardkey"`
	SmID              int         `json:"sm_id"`
	DateCreated       string      `json:"date_created"`
	OofShard          string      `json:"oof_shard"`
}

type deliveryDTO struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type paymentDTO struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type itemDTO struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func OrderFromJSON(data []byte) (entity.Order, error) {
	var dto orderDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, err
	}

	orderID, err := valueobject.NewOrderID(dto.OrderUID)
	if err != nil {
		return nil, err
	}
	trackNumber, err := valueobject.NewTrackNumber(dto.TrackNumber)
	if err != nil {
		return nil, err
	}
	entry, err := valueobject.NewOrderEntry(dto.Entry)
	if err != nil {
		return nil, err
	}
	locale, err := valueobject.NewLocale(dto.Locale)
	if err != nil {
		return nil, err
	}
	deliveryService, err := valueobject.NewDeliveryService(dto.DeliveryService)
	if err != nil {
		return nil, err
	}
	dateCreated, err := time.Parse(time.RFC3339, dto.DateCreated)
	if err != nil {
		return nil, err
	}

	// Delivery
	phone, err := valueobject.NewDeliveryPhone(dto.Delivery.Phone)
	if err != nil {
		return nil, err
	}
	email, err := valueobject.NewDeliveryEmail(dto.Delivery.Email)
	if err != nil {
		return nil, err
	}
	delivery := entity.NewDeliveryBuilder().
		SetName(dto.Delivery.Name).
		SetPhone(phone).
		SetZip(dto.Delivery.Zip).
		SetCity(dto.Delivery.City).
		SetAddress(dto.Delivery.Address).
		SetRegion(dto.Delivery.Region).
		SetEmail(email).
		BuildMust()

	// Payment
	currency, err := valueobject.NewCurrency(dto.Payment.Currency)
	if err != nil {
		return nil, err
	}
	provider, err := valueobject.NewPaymentProvider(dto.Payment.Provider)
	if err != nil {
		return nil, err
	}
	amount, _ := valueobject.NewAmount(dto.Payment.Amount)
	deliveryCost, _ := valueobject.NewAmount(dto.Payment.DeliveryCost)
	goodsTotal, _ := valueobject.NewAmount(dto.Payment.GoodsTotal)
	customFee, _ := valueobject.NewAmount(dto.Payment.CustomFee)
	payment := entity.NewPaymentBuilder().
		SetTransaction(dto.Payment.Transaction).
		SetRequestID(dto.Payment.RequestID).
		SetCurrency(currency).
		SetProvider(provider).
		SetAmount(amount).
		SetPaymentDT(dto.Payment.PaymentDT).
		SetBank(dto.Payment.Bank).
		SetDeliveryCost(deliveryCost).
		SetGoodsTotal(goodsTotal).
		SetCustomFee(customFee).
		BuildMust()

	// Items
	var items []entity.Item
	for _, i := range dto.Items {
		trackNum, _ := valueobject.NewTrackNumber(i.TrackNumber)
		price, _ := valueobject.NewAmount(i.Price)
		totalPrice, _ := valueobject.NewAmount(i.TotalPrice)
		rid, _ := valueobject.NewItemRID(i.RID)
		item := entity.NewItemBuilder().
			SetChrtID(i.ChrtID).
			SetTrackNumber(trackNum).
			SetPrice(price).
			SetRID(rid).
			SetName(i.Name).
			SetSale(i.Sale).
			SetSize(i.Size).
			SetTotalPrice(totalPrice).
			SetNmID(i.NmID).
			SetBrand(i.Brand).
			SetStatus(i.Status).
			BuildMust()
		items = append(items, item)
	}

	// Order
	return entity.NewOrderBuilder().
		SetUID(orderID).
		SetTrackNumber(trackNumber).
		SetEntry(entry).
		SetDelivery(delivery).
		SetPayment(payment).
		SetItems(items).
		SetLocale(locale).
		SetInternalSignature(dto.InternalSignature).
		SetCustomerID(dto.CustomerID).
		SetDeliveryService(deliveryService).
		SetShardKey(dto.ShardKey).
		SetSmID(dto.SmID).
		SetDateCreated(dateCreated).
		SetOofShard(dto.OofShard).
		Build()
}
