package services

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func (cfs *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {

	resp, err := http.Get(cfs.url)

	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()
	fact := &CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}
	return fact,nil
}

func NewCatFactService(url string) Service {

	return &CatFactService{

		url: url,
	}
}
