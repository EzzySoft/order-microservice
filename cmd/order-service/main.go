package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"order-service/internal/order/application"
	"order-service/internal/order/application/api"
	"order-service/internal/order/infrastructure/db/repository"
	"order-service/internal/shared/infrastructure/cache"
	"order-service/internal/shared/infrastructure/config"
	"order-service/internal/shared/infrastructure/db"
	"order-service/internal/shared/infrastructure/kafka"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// --- Конфиги и подключения к инфраструктуре ---

	redisCfg, err := config.LoadRedis(ctx)
	if err != nil {
		log.Fatalf("redis cfg: %v", err)
	}
	dbCfg, err := config.LoadDB(ctx)
	if err != nil {
		log.Fatalf("db cfg: %v", err)
	}
	kafkaCfg, err := config.LoadKafka(ctx)
	if err != nil {
		log.Fatalf("kafka cfg: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         redisCfg.Addr(),
		Password:     redisCfg.Password,
		DB:           redisCfg.DB,
		DialTimeout:  200 * time.Millisecond,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
	})
	dbConn, err := db.NewPostgres(dbCfg.DSN())
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	// --- Репозитории, сервисы, кэш ---

	orderRepo := repository.NewOrderSQLRepository(dbConn)
	orderUIDs, err := orderRepo.AllOrderUIDs(ctx)
	if err != nil {
		log.Fatalf("all order uids: %v", err)
	}
	log.Printf("[database] %d orders found from DB", len(orderUIDs))

	orderCache := cache.NewOrderCache(redisClient, orderRepo)
	orderCache.WarmUpFromRedis(ctx, orderUIDs)

	orderService := &application.OrderService{Repo: orderCache}
	orderAPI := &api.OrderAPI{Service: orderService}

	// --- HTTP server (статик + API) ---

	// 1. Фронтенд (отдача /web/index.html, /web/app.js)
	fs := http.FileServer(http.Dir(filepath.Join(".", "web")))

	mux := http.NewServeMux()
	mux.Handle("/", fs)             // Отдача статики по / и всему, что не /order/*
	mux.Handle("/order/", orderAPI) // Получить заказ по UID
	mux.Handle("/orders", orderAPI) // Получить список UID заказов

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	go func() {
		log.Printf("[API] HTTP server started on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// --- Kafka Consumer (фоново) ---

	kafkaConsumer := kafka.NewConsumer(
		kafkaCfg.Brokers, kafkaCfg.Topic, kafkaCfg.GroupID,
		orderService.HandleOrderJSON,
	)
	kafkaConsumer.Start(ctx)

	// --- Блокировка главной горутины ---
	select {}
}
