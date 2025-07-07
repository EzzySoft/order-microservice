package entity

import (
	"errors"
	"order-service/internal/order/domain/valueobject"
	"time"
)

type order struct {
	uid               valueobject.OrderID
	trackNumber       valueobject.TrackNumber
	entry             valueobject.OrderEntry
	delivery          Delivery
	payment           Payment
	items             []Item
	locale            valueobject.Locale
	internalSignature string
	customerID        string
	deliveryService   valueobject.DeliveryService
	shardKey          string
	smID              int
	dateCreated       time.Time
	oofShard          string
}

type Order interface {
	UID() valueobject.OrderID
	TrackNumber() valueobject.TrackNumber
	Entry() valueobject.OrderEntry
	Delivery() Delivery
	Payment() Payment
	Items() []Item
	Locale() valueobject.Locale
	InternalSignature() string
	CustomerID() string
	DeliveryService() valueobject.DeliveryService
	ShardKey() string
	SmID() int
	DateCreated() time.Time
	OofShard() string
}

func (o *order) UID() valueobject.OrderID                     { return o.uid }
func (o *order) TrackNumber() valueobject.TrackNumber         { return o.trackNumber }
func (o *order) Entry() valueobject.OrderEntry                { return o.entry }
func (o *order) Delivery() Delivery                           { return o.delivery }
func (o *order) Payment() Payment                             { return o.payment }
func (o *order) Items() []Item                                { return o.items }
func (o *order) Locale() valueobject.Locale                   { return o.locale }
func (o *order) InternalSignature() string                    { return o.internalSignature }
func (o *order) CustomerID() string                           { return o.customerID }
func (o *order) DeliveryService() valueobject.DeliveryService { return o.deliveryService }
func (o *order) ShardKey() string                             { return o.shardKey }
func (o *order) SmID() int                                    { return o.smID }
func (o *order) DateCreated() time.Time                       { return o.dateCreated }
func (o *order) OofShard() string                             { return o.oofShard }

type OrderBuilder struct {
	o *order
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{o: &order{}}
}

func (b *OrderBuilder) SetUID(uid valueobject.OrderID) *OrderBuilder {
	b.o.uid = uid
	return b
}
func (b *OrderBuilder) SetTrackNumber(tn valueobject.TrackNumber) *OrderBuilder {
	b.o.trackNumber = tn
	return b
}
func (b *OrderBuilder) SetEntry(e valueobject.OrderEntry) *OrderBuilder {
	b.o.entry = e
	return b
}
func (b *OrderBuilder) SetDelivery(d Delivery) *OrderBuilder {
	b.o.delivery = d
	return b
}
func (b *OrderBuilder) SetPayment(p Payment) *OrderBuilder {
	b.o.payment = p
	return b
}
func (b *OrderBuilder) SetItems(items []Item) *OrderBuilder {
	b.o.items = items
	return b
}
func (b *OrderBuilder) SetLocale(l valueobject.Locale) *OrderBuilder {
	b.o.locale = l
	return b
}
func (b *OrderBuilder) SetInternalSignature(s string) *OrderBuilder {
	b.o.internalSignature = s
	return b
}
func (b *OrderBuilder) SetCustomerID(c string) *OrderBuilder {
	b.o.customerID = c
	return b
}
func (b *OrderBuilder) SetDeliveryService(ds valueobject.DeliveryService) *OrderBuilder {
	b.o.deliveryService = ds
	return b
}
func (b *OrderBuilder) SetShardKey(sk string) *OrderBuilder {
	b.o.shardKey = sk
	return b
}
func (b *OrderBuilder) SetSmID(id int) *OrderBuilder {
	b.o.smID = id
	return b
}
func (b *OrderBuilder) SetDateCreated(dt time.Time) *OrderBuilder {
	b.o.dateCreated = dt
	return b
}
func (b *OrderBuilder) SetOofShard(oof string) *OrderBuilder {
	b.o.oofShard = oof
	return b
}

func (b *OrderBuilder) Build() (Order, error) {
	if b.o.uid == "" || b.o.trackNumber == "" || b.o.entry == "" || b.o.delivery == nil || b.o.payment == nil {
		return nil, errors.New("missing required fields")
	}
	return b.o, nil
}

func (b *OrderBuilder) BuildMust() Order {
	o, err := b.Build()
	if err != nil {
		panic(err)
	}
	return o
}
