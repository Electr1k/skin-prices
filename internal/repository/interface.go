package repository

import "skin-prices/internal/domain"

type PriceRepository interface {
	GetPrices() ([]*domain.Price, error)
	UpdateOrCreate(price *domain.Price) (*domain.Price, error)
}
