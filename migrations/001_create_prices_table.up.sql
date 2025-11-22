CREATE TABLE IF NOT EXISTS prices (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255),
    date DATE DEFAULT CURRENT_DATE NOT NULL,
    last_24h DECIMAL(10,2) NULL,
    last_7d DECIMAL(10,2) NULL,
    last_30d DECIMAL(10,2) NULL,
    last_90d DECIMAL(10,2) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name, date)
);

CREATE INDEX IF NOT EXISTS idx_prices_name ON prices(name);
CREATE INDEX IF NOT EXISTS idx_prices_date_name ON prices(date, name);

comment on column prices.id is 'Идентификатор';
comment on column prices.name is 'Название скина';
comment on column prices.date is 'Дата';
comment on column prices.last_24h is 'Медианная цена за последние 24 часа';
comment on column prices.last_7d is 'Медианная цена за последние 7 дней';
comment on column prices.last_30d is 'Медианная цена за последние 30 дней';
comment on column prices.last_90d is 'Медианная цена за последние 90 дней';