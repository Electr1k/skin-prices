package repository

import "awesomeProject/internal/domain"

type PriceRepository interface {
	GetPrices() ([]*domain.Price, error)
	UpdateOrCreate(price *domain.Price) (*domain.Price, error)
}
