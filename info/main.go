package main

import (
	"time"

	"github.com/gravision/alphaquark-upbit-api/etherscan"
	"github.com/gravision/alphaquark-upbit-api/exim"
	"github.com/gravision/alphaquark-upbit-api/goredis"
	"github.com/gravision/alphaquark-upbit-api/logger"
	"github.com/gravision/alphaquark-upbit-api/upbit"
	myutil "github.com/gravision/alphaquark-upbit-api/utils"
)

const (
	// Response
	RESP_SYMBOL     = "AQT"
	RESP_MAX_SUPPLY = 30000000
	RESP_PROVIDER   = "Alpha Quark Limited"
)

type Info struct {
	Symbol               string  `redis:"Symbol"`
	CurrencyCode         string  `redis:"CurrencyCode"`
	Price                float64 `redis:"Price"`
	MarketCap            uint64  `redis:"MarketCap"`
	AccTradePrice24H     string  `redis:"AccTradePrice24H"` // float
	CirculatingSupply    uint64  `redis:"CirculatingSupply"`
	MaxSupply            uint64  `redis:"MaxSupply"`
	Provider             string  `redis:"Provider"`
	LastUpdatedTimestamp int64   `redis:"LastUpdatedTimestamp"`
}

func main() {
	logger.Info("Start alphaquark api server")

	go func() {
		timer := time.NewTicker(time.Minute * 1)

		for {
			UpdatePrice()
			logger.Info("updated upbit")

			<-timer.C
		}
	}()

	go func() {
		timer := time.NewTicker(time.Hour * 6)

		for {
			UpdateExim()
			UpdateSupply()
			logger.Info("updated etherscan, exim")

			<-timer.C
		}
	}()

	go func() {
		for {
			timer := time.NewTicker(time.Minute * 1)
			<-timer.C

			data := CreateInfo()
			for k, v := range data {
				err := UpdateInfo(k, v)
				if err != nil {
					logger.Warn("cannot update information")
				}
			}
		}
	}()

	c := myutil.CleanupSignalHook()
	<-c
}

func UpdateSupply() {
	supply, err := etherscan.Serve()
	if err != nil {
		logger.Warn(err.Error())
		return
	}

	err = goredis.Insert("ETHERSCAN", supply)
	if err != nil {
		logger.Warn(err.Error())
	}
}

func UpdatePrice() {
	price, err := upbit.Serve()
	if err != nil {
		logger.Warn(err.Error())
		return
	}

	err = goredis.Insert("UPBIT", price)
	if err != nil {
		logger.Warn(err.Error())
	}
}

func UpdateExim() {
	exim, err := exim.Serve()
	if err != nil {
		logger.Warn(err.Error())
		return
	}

	goredis.HMSET("EXIM", exim)
	if err != nil {
		logger.Warn(err.Error())
		return
	}
}

func CreateInfo() (data map[string]Info) {
	tCirculatingSupply, err := goredis.GetUint("ETHERSCAN")
	if err != nil {
		logger.Warn(err.Error())

		return
	}

	tPrice, err := goredis.GetFloat64("UPBIT")
	if err != nil {
		logger.Warn(err.Error())

		return
	}

	var exchangeRate struct {
		IDR float64 `redis:"IDR"`
		USD float64 `redis:"USD"`
		SGD float64 `redis:"SGD"`
		THB float64 `redis:"THB"`
	}

	err = goredis.HGETALL("EXIM", &exchangeRate)
	if err != nil {
		logger.Warn(err.Error())
	}

	tTimestamp := time.Now().Unix()

	m := make(map[string]Info, 5) // KRW, USD, IDR, SGD, THB
	m["KRW"] = Info{
		Symbol:               RESP_SYMBOL,
		CurrencyCode:         "KRW",
		Price:                tPrice,
		MarketCap:            tCirculatingSupply * uint64(tPrice),
		AccTradePrice24H:     "null",
		CirculatingSupply:    tCirculatingSupply,
		MaxSupply:            RESP_MAX_SUPPLY,
		Provider:             RESP_PROVIDER,
		LastUpdatedTimestamp: tTimestamp,
	}

	tMap := make(map[string]float64)
	tMap["USD"] = exchangeRate.USD
	tMap["IDR"] = exchangeRate.IDR
	tMap["SGD"] = exchangeRate.SGD
	tMap["THB"] = exchangeRate.THB

	for k, v := range tMap {
		m[k] = Info{
			Symbol:               RESP_SYMBOL,
			CurrencyCode:         k,
			Price:                tPrice / v,
			MarketCap:            tCirculatingSupply * uint64(tPrice/v),
			AccTradePrice24H:     "null",
			CirculatingSupply:    tCirculatingSupply,
			MaxSupply:            RESP_MAX_SUPPLY,
			Provider:             RESP_PROVIDER,
			LastUpdatedTimestamp: tTimestamp,
		}
	}

	return m
}

func UpdateInfo(key string, data interface{}) (err error) {
	err = goredis.HMSET(key, data)

	return
}
