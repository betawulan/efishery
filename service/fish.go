package service

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"

	"github.com/betawulan/efishery/error_message"
	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/repository"
)

type fishService struct {
	fishRepo    repository.FishRepository
	UrlCurrency string
	SecretKey   []byte
}

func (f fishService) GetDataStorages(tokenString string) ([]model.Fish, error) {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return f.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, error_message.Unauthorized{Message: "token invalid"}
	}

	exchangeRateUSD, err := f.fishRepo.GetExchangeRate()
	if err != nil {
		return nil, err
	}

	fishes, err := f.fishRepo.GetFish()
	if err != nil {
		return nil, err
	}

	dollar := exchangeRateUSD.Data.USD
	for i, fish := range fishes {
		if fish.Price == "" {
			continue
		}

		rupiah, err := strconv.ParseFloat(fish.Price, 64)
		if err != nil {
			return nil, err
		}

		usd := rupiah * dollar
		usdString := fmt.Sprintf("%.3f", usd)
		fishes[i].PriceUSD = usdString
	}

	return fishes, nil
}

func NewFishService(fishRepo repository.FishRepository, urlCurrency string, secretKey []byte) FishService {
	return fishService{
		fishRepo:    fishRepo,
		UrlCurrency: urlCurrency,
		SecretKey:   secretKey,
	}
}
