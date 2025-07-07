package api

import (
	"encoding/json"
	"net/http"
	"order-service/internal/order/application"
	"order-service/internal/order/application/api/serializer"
	"strings"
)

type OrderAPI struct {
	Service *application.OrderService
}

func (h *OrderAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	if r.URL.Path == "/orders" {
		uids, err := h.Service.Repo.AllOrderUIDs(r.Context())
		if err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(uids)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/order/") {
		orderUID := strings.TrimPrefix(r.URL.Path, "/order/")
		if orderUID == "" {
			http.Error(w, "order_uid is required", http.StatusBadRequest)
			return
		}
		order, err := h.Service.Repo.FindByID(r.Context(), orderUID)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp := serializer.OrderToResponse(order)
		json.NewEncoder(w).Encode(resp)
		return
	}
	http.NotFound(w, r)
}
