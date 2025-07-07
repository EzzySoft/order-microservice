package mapper

import (
	"fmt"
	"order-service/internal/order/domain/entity"
	"order-service/internal/order/domain/valueobject"
	"order-service/internal/order/infrastructure/db/model"
)

// domain → model
func OrderEntityToModel(order entity.Order) (*model.Order, *model.Delivery, *model.Payment, []*model.Item) {
	o := &model.Order{
		UID:               string(order.UID()),
		TrackNumber:       string(order.TrackNumber()),
		Entry:             string(order.Entry()),
		InternalSignature: order.InternalSignature(),
		CustomerID:        order.CustomerID(),
		DeliveryService:   string(order.DeliveryService()),
		ShardKey:          order.ShardKey(),
		SmID:              order.SmID(),
		DateCreated:       order.DateCreated(),
		OofShard:          order.OofShard(),
	}
	delivery := order.Delivery()
	d := &model.Delivery{
		OrderUID: o.UID,
		Name:     delivery.Name(),
		Phone:    string(delivery.Phone()),
		Zip:      delivery.Zip(),
		City:     delivery.City(),
		Address:  delivery.Address(),
		Region:   delivery.Region(),
		Email:    string(delivery.Email()),
	}
	payment := order.Payment()
	p := &model.Payment{
		OrderUID:     o.UID,
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
	var items []*model.Item
	for _, it := range order.Items() {
		items = append(items, &model.Item{
			OrderUID:    o.UID,
			ChrtID:      it.ChrtID(),
			TrackNumber: string(it.TrackNumber()),
			Price:       it.Price().Int(),
			RID:         string(it.RID()),
			Name:        it.Name(),
			Sale:        it.Sale(),
			Size:        it.Size(),
			TotalPrice:  it.TotalPrice().Int(),
			NmID:        it.NmID(),
			Brand:       it.Brand(),
			Status:      it.Status(),
		})
	}
	return o, d, p, items
}

// model → domain
func ModelToOrderEntity(
	o model.Order,
	d model.Delivery,
	p model.Payment,
	items []model.Item,
) (entity.Order, error) {
	uid, err := valueobject.NewOrderID(o.UID)
	if err != nil {
		return nil, fmt.Errorf("invalid UID: %w", err)
	}
	trackNumber, err := valueobject.NewTrackNumber(o.TrackNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid TrackNumber: %w", err)
	}
	entry, err := valueobject.NewOrderEntry(o.Entry)
	if err != nil {
		return nil, fmt.Errorf("invalid Entry: %w", err)
	}
	deliveryService, err := valueobject.NewDeliveryService(o.DeliveryService)
	if err != nil {
		return nil, fmt.Errorf("invalid DeliveryService: %w", err)
	}
	locale := valueobject.LocaleRU // или вытащи из o если добавишь поле

	// delivery
	phone, err := valueobject.NewDeliveryPhone(d.Phone)
	if err != nil {
		return nil, fmt.Errorf("invalid DeliveryPhone: %w", err)
	}
	email, err := valueobject.NewDeliveryEmail(d.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid DeliveryEmail: %w", err)
	}
	delivery := entity.NewDeliveryBuilder().
		SetName(d.Name).
		SetPhone(phone).
		SetZip(d.Zip).
		SetCity(d.City).
		SetAddress(d.Address).
		SetRegion(d.Region).
		SetEmail(email).
		BuildMust()

	// payment
	currency, err := valueobject.NewCurrency(p.Currency)
	if err != nil {
		return nil, fmt.Errorf("invalid Currency: %w", err)
	}
	provider, err := valueobject.NewPaymentProvider(p.Provider)
	if err != nil {
		return nil, fmt.Errorf("invalid Provider: %w", err)
	}
	amount, err := valueobject.NewAmount(p.Amount)
	if err != nil {
		return nil, fmt.Errorf("invalid Amount: %w", err)
	}
	deliveryCost, err := valueobject.NewAmount(p.DeliveryCost)
	if err != nil {
		return nil, fmt.Errorf("invalid DeliveryCost: %w", err)
	}
	goodsTotal, err := valueobject.NewAmount(p.GoodsTotal)
	if err != nil {
		return nil, fmt.Errorf("invalid GoodsTotal: %w", err)
	}
	customFee, err := valueobject.NewAmount(p.CustomFee)
	if err != nil {
		return nil, fmt.Errorf("invalid CustomFee: %w", err)
	}
	payment := entity.NewPaymentBuilder().
		SetTransaction(p.Transaction).
		SetRequestID(p.RequestID).
		SetCurrency(currency).
		SetProvider(provider).
		SetAmount(amount).
		SetPaymentDT(p.PaymentDT).
		SetBank(p.Bank).
		SetDeliveryCost(deliveryCost).
		SetGoodsTotal(goodsTotal).
		SetCustomFee(customFee).
		BuildMust()

	// items
	var entityItems []entity.Item
	for _, it := range items {
		priceVO, err := valueobject.NewAmount(it.Price)
		if err != nil {
			return nil, fmt.Errorf("invalid Item.Price: %w", err)
		}
		totalPriceVO, err := valueobject.NewAmount(it.TotalPrice)
		if err != nil {
			return nil, fmt.Errorf("invalid Item.TotalPrice: %w", err)
		}
		trackNumberVO, err := valueobject.NewTrackNumber(it.TrackNumber)
		if err != nil {
			return nil, fmt.Errorf("invalid Item.TrackNumber: %w", err)
		}
		ridVO, err := valueobject.NewItemRID(it.RID)
		if err != nil {
			return nil, fmt.Errorf("invalid Item.RID: %w", err)
		}
		entityItems = append(entityItems, entity.NewItemBuilder().
			SetChrtID(it.ChrtID).
			SetTrackNumber(trackNumberVO).
			SetPrice(priceVO).
			SetRID(ridVO).
			SetName(it.Name).
			SetSale(it.Sale).
			SetSize(it.Size).
			SetTotalPrice(totalPriceVO).
			SetNmID(it.NmID).
			SetBrand(it.Brand).
			SetStatus(it.Status).
			BuildMust())
	}

	return entity.NewOrderBuilder().
		SetUID(uid).
		SetTrackNumber(trackNumber).
		SetEntry(entry).
		SetDelivery(delivery).
		SetPayment(payment).
		SetItems(entityItems).
		SetLocale(locale).
		SetInternalSignature(o.InternalSignature).
		SetCustomerID(o.CustomerID).
		SetDeliveryService(deliveryService).
		SetShardKey(o.ShardKey).
		SetSmID(o.SmID).
		SetDateCreated(o.DateCreated).
		SetOofShard(o.OofShard).
		Build()
}
