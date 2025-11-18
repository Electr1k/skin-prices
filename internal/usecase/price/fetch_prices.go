package price

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/repository"
	"awesomeProject/pkg/steam_data"
)

type FetchPricesUseCase struct {
	httpClient steam_data.PriceProvider
	repository repository.PriceRepository
}

func NewFetchPricesUseCase(
	httpClient steam_data.PriceProvider,
	repository repository.PriceRepository,
) *FetchPricesUseCase {
	return &FetchPricesUseCase{
		httpClient: httpClient,
		repository: repository,
	}
}

func (uc *FetchPricesUseCase) Handle() ([]*domain.Price, error) {
	data, err := uc.httpClient.FetchPrices()

	if err != nil {
		return nil, err
	}

	var prices []*domain.Price

	for skinName, dto := range data {

		price, err := domain.NewPrice(
			skinName,
			&dto.Steam.Last24h,
			&dto.Steam.Last7d,
			&dto.Steam.Last30d,
			&dto.Steam.Last90d,
		)

		if err != nil {
			return nil, err
		}

		model, err := uc.repository.UpdateOrCreate(price)
		if err != nil {
			return nil, err
		}

		prices = append(prices, model)
	}

	return prices, nil
}
