package entity

import (
	"errors"
	"order-service/internal/order/domain/valueobject"
)

type item struct {
	chrtID      int
	trackNumber valueobject.TrackNumber
	price       valueobject.Amount
	rid         valueobject.ItemRID
	name        string
	sale        int
	size        string
	totalPrice  valueobject.Amount
	nmID        int
	brand       string
	status      int
}

type Item interface {
	ChrtID() int
	TrackNumber() valueobject.TrackNumber
	Price() valueobject.Amount
	RID() valueobject.ItemRID
	Name() string
	Sale() int
	Size() string
	TotalPrice() valueobject.Amount
	NmID() int
	Brand() string
	Status() int
}

func (i *item) ChrtID() int                          { return i.chrtID }
func (i *item) TrackNumber() valueobject.TrackNumber { return i.trackNumber }
func (i *item) Price() valueobject.Amount            { return i.price }
func (i *item) RID() valueobject.ItemRID             { return i.rid }
func (i *item) Name() string                         { return i.name }
func (i *item) Sale() int                            { return i.sale }
func (i *item) Size() string                         { return i.size }
func (i *item) TotalPrice() valueobject.Amount       { return i.totalPrice }
func (i *item) NmID() int                            { return i.nmID }
func (i *item) Brand() string                        { return i.brand }
func (i *item) Status() int                          { return i.status }

type ItemBuilder struct {
	it *item
}

func NewItemBuilder() *ItemBuilder {
	return &ItemBuilder{it: &item{}}
}

func (b *ItemBuilder) SetChrtID(chrtID int) *ItemBuilder {
	b.it.chrtID = chrtID
	return b
}
func (b *ItemBuilder) SetTrackNumber(tn valueobject.TrackNumber) *ItemBuilder {
	b.it.trackNumber = tn
	return b
}
func (b *ItemBuilder) SetPrice(price valueobject.Amount) *ItemBuilder {
	b.it.price = price
	return b
}
func (b *ItemBuilder) SetRID(rid valueobject.ItemRID) *ItemBuilder {
	b.it.rid = rid
	return b
}
func (b *ItemBuilder) SetName(name string) *ItemBuilder {
	b.it.name = name
	return b
}
func (b *ItemBuilder) SetSale(sale int) *ItemBuilder {
	b.it.sale = sale
	return b
}
func (b *ItemBuilder) SetSize(size string) *ItemBuilder {
	b.it.size = size
	return b
}
func (b *ItemBuilder) SetTotalPrice(tp valueobject.Amount) *ItemBuilder {
	b.it.totalPrice = tp
	return b
}
func (b *ItemBuilder) SetNmID(nmID int) *ItemBuilder {
	b.it.nmID = nmID
	return b
}
func (b *ItemBuilder) SetBrand(brand string) *ItemBuilder {
	b.it.brand = brand
	return b
}
func (b *ItemBuilder) SetStatus(status int) *ItemBuilder {
	b.it.status = status
	return b
}

func (b *ItemBuilder) Build() (Item, error) {
	if b.it.trackNumber == "" || b.it.price < 0 {
		return nil, errors.New("missing or invalid required fields")
	}
	return b.it, nil
}

func (b *ItemBuilder) BuildMust() Item {
	i, err := b.Build()
	if err != nil {
		panic(err)
	}
	return i
}
