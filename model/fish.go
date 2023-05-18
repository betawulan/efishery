package model

import "time"

type Fish struct {
	Uuid         string    `json:"uuid"`
	Komoditas    string    `json:"komoditas"`
	AreaProvinsi string    `json:"area_provinsi"`
	AreaKota     string    `json:"area_kota"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	TglParsed    time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
	PriceUSD     string    `json:"price_usd"`
}

type ExchangeRate struct {
	USD float64 `json:"USD"`
}

type Currency struct {
	Data ExchangeRate `json:"data"`
}

type Summary struct {
	AreaProvinsi string  `json:"area_provinsi"`
	Max          int     `json:"max"`
	Min          int     `json:"min"`
	Avg          int     `json:"avg"`
	Median       float64 `json:"median"`
}
