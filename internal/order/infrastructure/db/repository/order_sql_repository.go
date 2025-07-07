package repository

import (
	"context"
	"order-service/internal/order/domain/entity"
	"order-service/internal/order/infrastructure/db/model"
	"order-service/internal/order/infrastructure/mapper"

	"github.com/jmoiron/sqlx"
)

type OrderSQLRepository struct {
	db *sqlx.DB
}

func NewOrderSQLRepository(db *sqlx.DB) *OrderSQLRepository {
	return &OrderSQLRepository{db: db}
}

func (r *OrderSQLRepository) Save(ctx context.Context, order entity.Order) error {
	o, d, p, items := mapper.OrderEntityToModel(order)
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// orders
	_, err = tx.NamedExec(`
		INSERT INTO orders (
			uid, track_number, entry, internal_signature, customer_id,
			delivery_service, shardkey, sm_id, date_created, oof_shard
		) VALUES (
			:uid, :track_number, :entry, :internal_signature, :customer_id,
			:delivery_service, :shardkey, :sm_id, :date_created, :oof_shard
		) ON CONFLICT(uid) DO NOTHING
	`, o)
	if err != nil {
		return err
	}

	// deliveries
	_, err = tx.NamedExec(`
		INSERT INTO deliveries (
			order_uid, name, phone, zip, city, address, region, email
		) VALUES (
			:order_uid, :name, :phone, :zip, :city, :address, :region, :email
		) ON CONFLICT(order_uid) DO NOTHING
	`, d)
	if err != nil {
		return err
	}

	// payments
	_, err = tx.NamedExec(`
		INSERT INTO payments (
			order_uid, transaction, request_id, currency, provider,
			amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
		) VALUES (
			:order_uid, :transaction, :request_id, :currency, :provider,
			:amount, :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee
		) ON CONFLICT(order_uid) DO NOTHING
	`, p)
	if err != nil {
		return err
	}

	// items (batch insert)
	for _, item := range items {
		_, err = tx.NamedExec(`
			INSERT INTO items (
				order_uid, chrt_id, track_number, price, rid, name, sale,
				size, total_price, nm_id, brand, status
			) VALUES (
				:order_uid, :chrt_id, :track_number, :price, :rid, :name, :sale,
				:size, :total_price, :nm_id, :brand, :status
			) ON CONFLICT DO NOTHING
		`, item)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderSQLRepository) FindByID(ctx context.Context, id string) (entity.Order, error) {
	var o model.Order
	if err := r.db.GetContext(ctx, &o, "SELECT * FROM orders WHERE uid=$1", id); err != nil {
		return nil, err
	}

	var d model.Delivery
	if err := r.db.GetContext(ctx, &d, "SELECT * FROM deliveries WHERE order_uid=$1", id); err != nil {
		return nil, err
	}

	var p model.Payment
	if err := r.db.GetContext(ctx, &p, "SELECT * FROM payments WHERE order_uid=$1", id); err != nil {
		return nil, err
	}

	var items []model.Item
	if err := r.db.SelectContext(ctx, &items, "SELECT * FROM items WHERE order_uid=$1", id); err != nil {
		return nil, err
	}

	return mapper.ModelToOrderEntity(o, d, p, items)
}

func (r *OrderSQLRepository) AllOrderUIDs(ctx context.Context) ([]string, error) {
	var ids []string
	err := r.db.SelectContext(ctx, &ids, "SELECT uid FROM orders")
	return ids, err
}
