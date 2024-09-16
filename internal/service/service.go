package service

import (
	"encoding/json"
	"net/http"
)

type Service struct{
	client *http.Client
	url string
}

func New(_url string) *Service{
	_client := &http.Client{}
	
	return &Service{
		client: _client,
		url: _url,
	}
}

type SupplyResponse struct {
    Supply []SupplyItem `json:"supply"`
}

type SupplyItem struct {
    Amount string `json:"amount"`
}

func (s *Service) GetAmount() (string, error){
	resp, err := s.client.Get(s.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jsonData SupplyResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonData); err != nil {
		return "", err
	}

	return jsonData.Supply[0].Amount, nil	
}