package postgres

import (
	"awesomeProject/internal/domain"
	"context"
	"database/sql"
	"fmt"
)

func (p *Postgres) GetPrices() ([]domain.Price, error) {
	rows, err := p.pool.Query(context.Background(), `
		SELECT name, last_24h, last_7d, last_30d, last_90d
		FROM prices;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var prices []domain.Price

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

		price, err := domain.NewPrice(
			name,
			nullFloatToFloat32Ptr(last24h),
			nullFloatToFloat32Ptr(last7d),
			nullFloatToFloat32Ptr(last30d),
			nullFloatToFloat32Ptr(last90d),
		)
		if err != nil {
			return nil, err
		}

		prices = append(prices, *price)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}

func nullFloatToFloat32Ptr(n sql.NullFloat64) *float32 {
	if n.Valid {
		val := float32(n.Float64)
		return &val
	}
	return nil
}
