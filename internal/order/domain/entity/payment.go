package entity

import (
	"errors"
	"order-service/internal/order/domain/valueobject"
)

type payment struct {
	transaction  string
	requestID    string
	currency     valueobject.Currency
	provider     valueobject.PaymentProvider
	amount       valueobject.Amount
	paymentDT    int64
	bank         string
	deliveryCost valueobject.Amount
	goodsTotal   valueobject.Amount
	customFee    valueobject.Amount
}

type Payment interface {
	Transaction() string
	RequestID() string
	Currency() valueobject.Currency
	Provider() valueobject.PaymentProvider
	Amount() valueobject.Amount
	PaymentDT() int64
	Bank() string
	DeliveryCost() valueobject.Amount
	GoodsTotal() valueobject.Amount
	CustomFee() valueobject.Amount
}

func (p *payment) Transaction() string                   { return p.transaction }
func (p *payment) RequestID() string                     { return p.requestID }
func (p *payment) Currency() valueobject.Currency        { return p.currency }
func (p *payment) Provider() valueobject.PaymentProvider { return p.provider }
func (p *payment) Amount() valueobject.Amount            { return p.amount }
func (p *payment) PaymentDT() int64                      { return p.paymentDT }
func (p *payment) Bank() string                          { return p.bank }
func (p *payment) DeliveryCost() valueobject.Amount      { return p.deliveryCost }
func (p *payment) GoodsTotal() valueobject.Amount        { return p.goodsTotal }
func (p *payment) CustomFee() valueobject.Amount         { return p.customFee }

type PaymentBuilder struct {
	p *payment
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{p: &payment{}}
}

func (b *PaymentBuilder) SetTransaction(t string) *PaymentBuilder {
	b.p.transaction = t
	return b
}
func (b *PaymentBuilder) SetRequestID(r string) *PaymentBuilder {
	b.p.requestID = r
	return b
}
func (b *PaymentBuilder) SetCurrency(c valueobject.Currency) *PaymentBuilder {
	b.p.currency = c
	return b
}
func (b *PaymentBuilder) SetProvider(pr valueobject.PaymentProvider) *PaymentBuilder {
	b.p.provider = pr
	return b
}
func (b *PaymentBuilder) SetAmount(a valueobject.Amount) *PaymentBuilder {
	b.p.amount = a
	return b
}
func (b *PaymentBuilder) SetPaymentDT(dt int64) *PaymentBuilder {
	b.p.paymentDT = dt
	return b
}
func (b *PaymentBuilder) SetBank(bk string) *PaymentBuilder {
	b.p.bank = bk
	return b
}
func (b *PaymentBuilder) SetDeliveryCost(ac valueobject.Amount) *PaymentBuilder {
	b.p.deliveryCost = ac
	return b
}
func (b *PaymentBuilder) SetGoodsTotal(gt valueobject.Amount) *PaymentBuilder {
	b.p.goodsTotal = gt
	return b
}
func (b *PaymentBuilder) SetCustomFee(cf valueobject.Amount) *PaymentBuilder {
	b.p.customFee = cf
	return b
}

func (b *PaymentBuilder) Build() (Payment, error) {
	if b.p.currency == "" || b.p.provider == "" {
		return nil, errors.New("missing required fields")
	}
	return b.p, nil
}

func (b *PaymentBuilder) BuildMust() Payment {
	p, err := b.Build()
	if err != nil {
		panic(err)
	}
	return p
}
