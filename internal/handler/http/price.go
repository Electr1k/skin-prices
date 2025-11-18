package http

import (
	"encoding/json"
	"net/http"
	"skin-prices/internal/usecase/price"
)

type PriceHandler struct {
	getPriceUseCase   *price.GetPricesUseCase
	fetchPriceUseCase *price.FetchPricesUseCase
}

func NewPriceHandler(
	getPriceUseCase *price.GetPricesUseCase,
	fetchPriceUseCase *price.FetchPricesUseCase,
) *PriceHandler {
	return &PriceHandler{
		getPriceUseCase:   getPriceUseCase,
		fetchPriceUseCase: fetchPriceUseCase,
	}
}

func (h *PriceHandler) GetPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := h.getPriceUseCase.Handle()
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
