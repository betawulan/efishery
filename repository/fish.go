package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/betawulan/efishery/model"
)

type fishRepo struct {
	UrlCurrency  string
	UrlFish      string
	exchangeRate float64
	date         string
}

func (f *fishRepo) GetExchangeRate() (model.Currency, error) {
	if f.date != "" && f.exchangeRate != 0 {
		if time.Now().Format("2006-01-02") == f.date {
			return model.Currency{Data: model.ExchangeRate{USD: f.exchangeRate}}, nil
		}
	}

	resp, err := http.Get(f.UrlCurrency)
	if err != nil {
		return model.Currency{}, err
	}

	defer func() {
		errClose := resp.Body.Close()
		if errClose != nil {
		}
	}()

	var currency model.Currency
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &currency)
	if err != nil {
		return model.Currency{}, err
	}

	// set
	date := time.Now().Format("2006-01-02")
	f.date = date
	f.exchangeRate = currency.Data.USD

	return currency, nil
}

func (f *fishRepo) GetFish() ([]model.Fish, error) {
	resp, err := http.Get(f.UrlFish)
	if err != nil {
		return []model.Fish{}, err
	}

	defer func() {
		errClose := resp.Body.Close()
		if errClose != nil {
		}
	}()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []model.Fish{}, err
	}

	fishes := make([]model.Fish, 0)
	err = json.Unmarshal(bodyBytes, &fishes)
	if err != nil {
		return []model.Fish{}, err
	}

	return fishes, nil
}

func NewFishRepository(urlCurrency, urlFish string) FishRepository {
	return &fishRepo{
		UrlCurrency: urlCurrency,
		UrlFish:     urlFish,
	}
}
