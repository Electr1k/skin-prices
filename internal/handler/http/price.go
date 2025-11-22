package http

import (
	"encoding/json"
	"net/http"
	"skin-prices/internal/dto"
	"skin-prices/internal/usecase/price"
	"time"
)

type PriceHandler struct {
	getPriceUseCase   *price.GetPricesUseCase
	fetchPriceUseCase *price.FetchPricesUseCase
}

type priceResponse struct {
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

	var priceDTOs []dto.PriceResponse

	for _, priceDomain := range prices {

		last24h := float32(*priceDomain.Last24h)
		last7d := float32(*priceDomain.Last7d)
		last30d := float32(*priceDomain.Last30d)
		last90d := float32(*priceDomain.Last90d)
		priceDTO, _ := dto.NewPriceResponse(string(priceDomain.Name), time.Time(priceDomain.Date), &last24h, &last7d, &last30d, &last90d)

		priceDTOs = append(priceDTOs, *priceDTO)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(priceDTOs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *PriceHandler) FetchPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := h.fetchPriceUseCase.Handle()

	var priceDTOs []dto.PriceResponse

	for _, priceDomain := range prices {

		last24h := float32(*priceDomain.Last24h)
		last7d := float32(*priceDomain.Last7d)
		last30d := float32(*priceDomain.Last30d)
		last90d := float32(*priceDomain.Last90d)
		priceDTO, _ := dto.NewPriceResponse(string(priceDomain.Name), time.Time(priceDomain.Date), &last24h, &last7d, &last30d, &last90d)

		priceDTOs = append(priceDTOs, *priceDTO)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(priceDTOs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
