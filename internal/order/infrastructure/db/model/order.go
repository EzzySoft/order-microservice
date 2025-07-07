package model

import "time"

type Order struct {
	UID               string    `db:"uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          string    `db:"shardkey"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}
