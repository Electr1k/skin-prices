package price

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
)

type GetPricesUseCase struct {
	priceRepo repository.PriceRepository
}

func NewGetPricesUseCase(priceRepo repository.PriceRepository) *GetPricesUseCase {
	return &GetPricesUseCase{
		priceRepo: priceRepo,
	}
}

func (uc *GetPricesUseCase) Handle() ([]models.Price, error) {
	return uc.priceRepo.GetPrices()
}
