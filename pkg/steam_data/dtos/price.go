package dtos

type PriceResponseDTO map[string]struct {
	Steam struct {
		Last24h float32 `json:"last_24h"`
		Last7d  float32 `json:"last_7d"`
		Last30d float32 `json:"last_30d"`
		Last90d float32 `json:"last_90d"`
	} `json:"steam"`
}
