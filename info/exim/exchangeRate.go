package exim

import (
	"strconv"
	"strings"

	"github.com/gravision/alphaquark-upbit-api/logger"
)

func (c *Client) ExchangeRate() (out map[string]float64, err error) {
	logger.Debug("Try to get exchange rate from exim")

	rates, err := c.call(exim_apiKey, "AP01")
	if err != nil {
		return
	}

	checkCurrency := []string{
		"USD", "IDR(100)", "SGD", "THB",
	}

	out = make(map[string]float64, 5)
	for i := 0; i < len(checkCurrency); i++ {
		for j := 0; j < len(rates); j++ {
			if rates[j].CUR_UNIT == checkCurrency[i] {
				price := strings.Replace(rates[j].DEAL_BAS_R, ",", "", -1)

				if checkCurrency[i] == "IDR(100)" {
					out["IDR"], _ = strconv.ParseFloat(price, 64)
				} else {
					out[checkCurrency[i]], _ = strconv.ParseFloat(price, 64)
				}

				logger.Info(checkCurrency[i] + " price : " + price)
				break
			}
		}
	}

	return
}
