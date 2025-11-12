package repository

import "awesomeProject/internal/domain"

type PriceRepository interface {
	GetPrices() ([]domain.Price, error)
}
