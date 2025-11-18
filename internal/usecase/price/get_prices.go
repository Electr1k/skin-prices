package price

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/repository"
)

type GetPricesUseCase struct {
	priceRepository repository.PriceRepository
}

func NewGetPricesUseCase(
	priceRepository repository.PriceRepository,
) *GetPricesUseCase {
	return &GetPricesUseCase{
		priceRepository: priceRepository,
	}
}

func (uc *GetPricesUseCase) Handle() ([]*domain.Price, error) {
	return uc.priceRepository.GetPrices()
}
