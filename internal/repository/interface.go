package repository

import "awesomeProject/internal/models"

type PriceRepository interface {
	GetPrices() ([]models.Price, error)
}
