package steam_data

import (
	"skin-prices/pkg/steam_data/dtos"
)

type PriceProvider interface {
	FetchPrices() (dtos.PriceResponseDTO, error)
}
