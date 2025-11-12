package http

import (
	"awesomeProject/internal/usecase/price"
	"encoding/json"
	"net/http"
)

type PriceHandler struct {
	priceUseCase *price.GetPricesUseCase
}

func NewPriceHandler(priceUseCase *price.GetPricesUseCase) *PriceHandler {
	return &PriceHandler{
		priceUseCase: priceUseCase,
	}
}

func (h *PriceHandler) GetPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := h.priceUseCase.Handle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prices); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
