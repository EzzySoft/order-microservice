package entity

import (
	"errors"
	"order-service/internal/order/domain/valueobject"
)

type delivery struct {
	name    string
	phone   valueobject.DeliveryPhone
	zip     string
	city    string
	address string
	region  string
	email   valueobject.DeliveryEmail
}

type Delivery interface {
	Name() string
	Phone() valueobject.DeliveryPhone
	Zip() string
	City() string
	Address() string
	Region() string
	Email() valueobject.DeliveryEmail
}

func (d *delivery) Name() string                     { return d.name }
func (d *delivery) Phone() valueobject.DeliveryPhone { return d.phone }
func (d *delivery) Zip() string                      { return d.zip }
func (d *delivery) City() string                     { return d.city }
func (d *delivery) Address() string                  { return d.address }
func (d *delivery) Region() string                   { return d.region }
func (d *delivery) Email() valueobject.DeliveryEmail { return d.email }

type DeliveryBuilder struct {
	d *delivery
}

func NewDeliveryBuilder() *DeliveryBuilder {
	return &DeliveryBuilder{d: &delivery{}}
}

func (b *DeliveryBuilder) SetName(name string) *DeliveryBuilder {
	b.d.name = name
	return b
}
func (b *DeliveryBuilder) SetPhone(phone valueobject.DeliveryPhone) *DeliveryBuilder {
	b.d.phone = phone
	return b
}
func (b *DeliveryBuilder) SetZip(zip string) *DeliveryBuilder {
	b.d.zip = zip
	return b
}
func (b *DeliveryBuilder) SetCity(city string) *DeliveryBuilder {
	b.d.city = city
	return b
}
func (b *DeliveryBuilder) SetAddress(address string) *DeliveryBuilder {
	b.d.address = address
	return b
}
func (b *DeliveryBuilder) SetRegion(region string) *DeliveryBuilder {
	b.d.region = region
	return b
}
func (b *DeliveryBuilder) SetEmail(email valueobject.DeliveryEmail) *DeliveryBuilder {
	b.d.email = email
	return b
}

func (b *DeliveryBuilder) Build() (Delivery, error) {
	if b.d.phone == "" || b.d.email == "" {
		return nil, errors.New("missing required fields")
	}
	return b.d, nil
}

func (b *DeliveryBuilder) BuildMust() Delivery {
	d, err := b.Build()
	if err != nil {
		panic(err)
	}
	return d
}
