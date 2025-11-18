package steam_data

import (
	"awesomeProject/pkg/steam_data/dtos"
)

type PriceProvider interface {
	FetchPrices() (dtos.PriceResponseDTO, error)
}
