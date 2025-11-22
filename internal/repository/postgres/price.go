package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"skin-prices/internal/domain"
	"time"
)

func (p *Postgres) GetPrices() ([]*domain.Price, error) {
	rows, err := p.pool.Query(context.Background(), `
		SELECT name, date, last_24h, last_7d, last_30d, last_90d
		FROM prices;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var prices []*domain.Price

	for rows.Next() {
		var (
			name                              string
			date                              time.Time
			last24h, last7d, last30d, last90d sql.NullFloat64
		)

		if err := rows.Scan(
			&name,
			&date,
			&last24h,
			&last7d,
			&last30d,
			&last90d,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		price, err := domain.NewPrice(
			name,
			date,
			nullFloatToFloat32Ptr(last24h),
			nullFloatToFloat32Ptr(last7d),
			nullFloatToFloat32Ptr(last30d),
			nullFloatToFloat32Ptr(last90d),
		)
		if err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}

func (p *Postgres) UpdateOrCreate(price *domain.Price) (*domain.Price, error) {

	row := p.pool.QueryRow(context.Background(),
		`
			INSERT INTO prices (name, date, last_24h, last_7d, last_30d, last_90d) 
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (name, date) 
			DO UPDATE SET 
				last_24h = EXCLUDED.last_24h,
				last_7d = EXCLUDED.last_7d,
				last_30d = EXCLUDED.last_30d,
				last_90d = EXCLUDED.last_90d,
				updated_at = CURRENT_TIMESTAMP
			RETURNING name, date, last_24h, last_7d, last_30d, last_90d;
		`, price.Name, price.Date, price.Last24h, price.Last7d, price.Last30d, price.Last90d)

	var (
		name                              string
		date                              time.Time
		last24h, last7d, last30d, last90d sql.NullFloat64
	)

	if err := row.Scan(
		&name,
		&date,
		&last24h,
		&last7d,
		&last30d,
		&last90d,
	); err != nil {
		return nil, err
	}

	dto, err := domain.NewPrice(
		name,
		date,
		nullFloatToFloat32Ptr(last24h),
		nullFloatToFloat32Ptr(last7d),
		nullFloatToFloat32Ptr(last30d),
		nullFloatToFloat32Ptr(last90d),
	)

	if err != nil {
		return nil, err
	}

	return dto, nil
}
