package cron

import (
	"log"
	"skin-prices/internal/usecase/price"
	"strconv"
)

type FetchPriceTask struct {
	fetchPriceUseCase *price.FetchPricesUseCase
}

func NewFetchPriceTask(
	fetchPriceUseCase *price.FetchPricesUseCase,
) *FetchPriceTask {
	return &FetchPriceTask{
		fetchPriceUseCase: fetchPriceUseCase,
	}
}

func (h *FetchPriceTask) Handle() {
	log.Println("Run cron job")
	items, err := h.fetchPriceUseCase.Handle()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Run cron job finished. Saved: " + strconv.Itoa(len(items)))
}
