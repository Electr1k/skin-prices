package postgres

import (
	"awesomeProject/internal/models"
	"context"
	"database/sql"
	"fmt"
)

func (p *Postgres) GetPrices() ([]models.Price, error) {
	rows, err := p.pool.Query(context.Background(), `
		SELECT name, last_24h, last_7d, last_30d, last_90d
		FROM prices;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query prices: %w", err)
	}
	defer rows.Close()

	var prices []models.Price
	var count int

	for rows.Next() {
		var (
			name                              string
			last24h, last7d, last30d, last90d sql.NullFloat64
		)

		if err := rows.Scan(
			&name,
			&last24h,
			&last7d,
			&last30d,
			&last90d,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		price, err := models.NewPrice(
			name,
			nullFloatToFloat32Ptr(last24h),
			nullFloatToFloat32Ptr(last7d),
			nullFloatToFloat32Ptr(last30d),
			nullFloatToFloat32Ptr(last90d),
		)
		if err != nil {
			fmt.Printf("Failed to create price models for %s: %v\n", name, err)
			continue
		}

		prices = append(prices, *price)
		count++
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	fmt.Printf("Successfully processed %d price records\n", count)
	return prices, nil
}

func nullFloatToFloat32Ptr(n sql.NullFloat64) *float32 {
	if n.Valid {
		val := float32(n.Float64)
		return &val
	}
	return nil
}
