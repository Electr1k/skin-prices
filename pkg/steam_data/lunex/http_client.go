package lunex

import (
	"awesomeProject/pkg/steam_data/dtos"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	baseURL string
	client  *http.Client
}

func NewClient() *HttpClient {
	return &HttpClient{
		baseURL: "https://raw.githubusercontent.com/LukeX404/",
		client:  &http.Client{Timeout: 60 * time.Second},
	}
}

func (c *HttpClient) makeRequest(uri string, method string) ([]byte, error) {
	req, err := http.NewRequest(method, c.baseURL+uri, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	response, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *HttpClient) FetchPrices() (dtos.PriceResponseDTO, error) {
	response, err := c.makeRequest("cs2-prices-tracker/refs/heads/main/static/prices/latest.json", "GET")

	if err != nil {
		panic(err)
		return nil, err
	}

	var priceResponse dtos.PriceResponseDTO
	if err := json.Unmarshal(response, &priceResponse); err != nil {
		return nil, err
	}

	return priceResponse, nil
}
