package domain

import "github.com/go-playground/validator/v10"

type SkinName string

type Dollar float32

var validate = validator.New(validator.WithRequiredStructEnabled())

type Price struct {
	Name    SkinName `json:"name" validate:"required"`
	Last24h *Dollar  `json:"last_24h" validate:"omitempty,min=0"`
	Last7d  *Dollar  `json:"last_7d" validate:"omitempty,min=0"`
	Last30d *Dollar  `json:"last_30d" validate:"omitempty,min=0"`
	Last90d *Dollar  `json:"last_90d" validate:"omitempty,min=0"`
}

func NewPrice(
	name string,
	last24h *float32,
	last7d *float32,
	last30d *float32,
	last90d *float32,
) (*Price, error) {
	p := &Price{
		Name: SkinName(name),
	}

	if last24h != nil {
		dollar := Dollar(*last24h)
		p.Last24h = &dollar
	}
	if last7d != nil {
		dollar := Dollar(*last7d)
		p.Last7d = &dollar
	}
	if last30d != nil {
		dollar := Dollar(*last30d)
		p.Last30d = &dollar
	}
	if last90d != nil {
		dollar := Dollar(*last90d)
		p.Last90d = &dollar
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Price) Validate() error {
	return validate.Struct(p)
}
