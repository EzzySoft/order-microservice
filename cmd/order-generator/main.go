package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

const (
	TotalOrders  = 100
	ProgressStep = 1
)

type Order struct {
	OrderUID          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Item   `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	ShardKey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
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

type Item struct {
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

func randomOrderUID() string {
	hex := uuid.New().String()
	hex = strings.ReplaceAll(hex, "-", "")
	return hex[:19]
}

func randomOrder() Order {
	orderUID := randomOrderUID()
	track := "WB" + uuid.NewString()[:8]
	now := time.Now().UTC().Format(time.RFC3339)

	phone := "+7"
	for i := 0; i < 10; i++ {
		phone += string('0' + rand.Intn(10))
	}
	email := uuid.NewString()[:8] + "@example.com"

	names := []string{"Vasya Pupkin", "Anna Ivanova", "Oleg Sidorov", "Igor Petrov"}
	cities := []string{"Moscow", "Spb", "Novosibirsk", "Ekaterinburg"}
	addresses := []string{"Lenina 1", "Karla Marksa 22", "Pushkina 3", "Tverskaya 7"}
	regions := []string{"Region1", "Region2", "Region3"}

	return Order{
		OrderUID:    orderUID,
		TrackNumber: track,
		Entry:       "WBIL",
		Delivery: Delivery{
			Name:    names[rand.Intn(len(names))],
			Phone:   phone,
			Zip:     "1" + string('0'+rand.Intn(9)) + "39809",
			City:    cities[rand.Intn(len(cities))],
			Address: addresses[rand.Intn(len(addresses))],
			Region:  regions[rand.Intn(len(regions))],
			Email:   email,
		},
		Payment: Payment{
			Transaction:  orderUID,
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(5000) + 1000,
			PaymentDT:    time.Now().Unix(),
			Bank:         "alpha",
			DeliveryCost: 1000 + rand.Intn(1000),
			GoodsTotal:   200 + rand.Intn(500),
			CustomFee:    rand.Intn(50),
		},
		Items: []Item{
			{
				ChrtID:      9934930 + rand.Intn(1000),
				TrackNumber: track,
				Price:       400 + rand.Intn(100),
				RID:         uuid.NewString(),
				Name:        "Item-" + uuid.NewString()[:4],
				Sale:        rand.Intn(50),
				Size:        string('A' + rand.Intn(5)),
				TotalPrice:  300 + rand.Intn(200),
				NmID:        2000000 + rand.Intn(100000),
				Brand:       "Brand-" + uuid.NewString()[:5],
				Status:      200 + rand.Intn(10),
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        uuid.NewString()[:10],
		DeliveryService:   "meest",
		ShardKey:          string('0' + rand.Intn(9)),
		SmID:              90 + rand.Intn(10),
		DateCreated:       now,
		OofShard:          string('0' + rand.Intn(5)),
	}
}

func printProgress(sent, total int) {
	percent := float64(sent) / float64(total) * 100
	barLen := 40
	filled := int(percent / 100 * float64(barLen))
	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := filled; i < barLen; i++ {
		bar += " "
	}
	fmt.Fprintf(os.Stdout, "\r[%s] %6.2f%% (%d/%d)", bar, percent, sent, total)
	if sent == total {
		fmt.Println()
	}
}

func main() {
	brokers := []string{"localhost:19092", "localhost:19093", "localhost:19094"}
	topic := "orders"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Async:   true,
		Dialer:  &kafka.Dialer{Timeout: 3 * time.Second},
	})
	defer writer.Close()

	ctx := context.Background()

	for i := 1; i <= TotalOrders; i++ {
		order := randomOrder()
		data, err := json.Marshal(order)
		if err != nil {
			fmt.Fprintf(os.Stderr, "marshal error: %v\n", err)
			continue
		}
		msg := kafka.Message{
			Key:   []byte(order.OrderUID),
			Value: data,
		}
		if err := writer.WriteMessages(ctx, msg); err != nil {
			fmt.Fprintf(os.Stderr, "kafka write error: %v\n", err)
		}
		if i%ProgressStep == 0 || i == TotalOrders {
			printProgress(i, TotalOrders)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
