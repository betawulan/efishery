package service

import (
	"context"
	"fmt"
	"sort"
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

func (f fishService) Summary(ctx context.Context, tokenString string) ([]model.Summary, error) {
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

	if claim.Role != "admin" {
		return nil, error_message.Unauthorized{Message: "you are not admin"}
	}

	fishes, err := f.fishRepo.GetFish()
	if err != nil {
		return nil, err
	}

	mapProvinsi := map[string]struct {
		isExist bool
		index   int
		price   []int
		sum     int
	}{}

	var index int
	var summaries []model.Summary
	for _, fish := range fishes {
		if fish.AreaProvinsi == "" {
			continue
		}

		if mapProvinsi[fish.AreaProvinsi].isExist {
			priceInt, err := strconv.Atoi(fish.Price)
			if err != nil {
				return nil, err
			}

			_sum := mapProvinsi[fish.AreaProvinsi].sum + priceInt
			allPrice := append(mapProvinsi[fish.AreaProvinsi].price, priceInt)

			mapProvinsi[fish.AreaProvinsi] = struct {
				isExist bool
				index   int
				price   []int
				sum     int
			}{isExist: mapProvinsi[fish.AreaProvinsi].isExist, index: mapProvinsi[fish.AreaProvinsi].index, price: allPrice, sum: _sum}
		} else {
			priceInt, err := strconv.Atoi(fish.Price)
			if err != nil {
				return nil, err
			}

			mapProvinsi[fish.AreaProvinsi] = struct {
				isExist bool
				index   int
				price   []int
				sum     int
			}{isExist: true, index: index, price: []int{priceInt}, sum: priceInt}
			index++
		}
	}

	for areaProvinsi, val := range mapProvinsi {
		sort.Ints(val.price)

		var summary model.Summary
		summary.AreaProvinsi = areaProvinsi

		lenPrice := len(val.price)
		if lenPrice > 0 {
			summary.Min = val.price[0]
			summary.Max = val.price[lenPrice-1]
			summary.Avg = float64(val.sum) / float64(lenPrice)

			if lenPrice%2 == 1 {
				indexMedian := (lenPrice / 2) + 1
				summary.Median = float64(val.price[indexMedian])
			} else {
				indexMedianFirst := (lenPrice / 2) - 1
				indexMedianSecond := (lenPrice / 2) + 1
				summary.Median = (float64(val.price[indexMedianFirst]) + float64(val.price[indexMedianSecond])) / 2
			}

		}

		summaries = append(summaries, summary)

	}

	return summaries, nil
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
