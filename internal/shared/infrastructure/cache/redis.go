package cache

import (
	"context"
	"encoding/json"
	"log"
	"order-service/internal/order/application/api/serializer"
	"order-service/internal/order/domain/entity"
	factory "order-service/internal/order/domain/factory"
	iface "order-service/internal/order/interface"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewOrderCache(redis *redis.Client, repo iface.OrderRepository) *OrderCache {
	return &OrderCache{
		redis: redis,
		repo:  repo,
	}
}

type OrderCache struct {
	redis *redis.Client
	repo  iface.OrderRepository
}

func (c *OrderCache) AllOrderUIDs(ctx context.Context) ([]string, error) {
	return c.repo.AllOrderUIDs(ctx)
}

func (c *OrderCache) Save(ctx context.Context, order entity.Order) error {
	id := string(order.UID())

	if err := c.repo.Save(ctx, order); err != nil {
		log.Printf("[cache] ❌ SAVE order %s to DB: %v", id, err)
		return err
	}

	dto := serializer.OrderToResponse(order)
	data, err := json.Marshal(dto)
	if err != nil {
		log.Printf("[cache] ❌ Marshal order %s: %v", id, err)
		return err
	}
	err = c.redis.Set(ctx, id, data, time.Hour).Err()
	if err == nil {
		log.Printf("[cache] ✅ SAVE order %s to Redis", id)
	} else {
		log.Printf("[cache] ❌ SAVE order %s to Redis ERROR: %v", id, err)
	}
	return err
}

func (c *OrderCache) Get(ctx context.Context, id string) (entity.Order, bool) {
	data, err := c.redis.Get(ctx, id).Bytes()
	if err == nil {
		log.Printf("[cache] ✅ HIT for %s", id)
		order, err := factory.OrderFromJSON(data)
		if err == nil {
			return order, true
		}
		log.Printf("[cache] ❌ Unmarshal order %s from Redis: %v", id, err)
	}
	log.Printf("[cache] 🚫 MISS for %s", id)
	return nil, false
}

func (c *OrderCache) WarmUpFromRedis(ctx context.Context, ids []string) {
	if err := c.redis.Ping(ctx).Err(); err != nil {
		log.Printf("[cache] 🚫 Redis unavailable! \n\n ERROR: %v\n\n Skipping...", err)
		return
	}

	for _, id := range ids {
		_, err := c.redis.Get(ctx, id).Bytes()
		if err == nil {
			log.Printf("[cache] 🔥 Already in Redis: %s", id)
			continue
		}
		order, err := c.repo.FindByID(ctx, id)
		if err == nil {
			data, err := json.Marshal(order)
			if err == nil {
				c.redis.Set(ctx, id, data, time.Hour)
				log.Printf("[cache] 🔥 Warmed up: %s", id)

			}
		}
	}
}

func (c *OrderCache) FindByID(ctx context.Context, id string) (entity.Order, error) {
	order, ok := c.Get(ctx, id)
	if ok {
		log.Printf("[cache] ▶️  RETURN FROM CACHE: %s", id)
		return order, nil
	}
	log.Printf("[cache] 🔍  FETCHING from DB... [%s]", id)
	order, err := c.repo.FindByID(ctx, id)
	if err != nil {
		log.Printf("[cache] ❌  DB lookup fail: %s, %v", id, err)
		return nil, err
	}
	c.Save(ctx, order)
	log.Printf("[cache] ▶️  RETURN FROM DB: %s", id)
	return order, nil
}
