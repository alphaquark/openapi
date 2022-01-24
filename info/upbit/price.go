package upbit

import (
	"fmt"

	"github.com/gravision/alphaquark-upbit-api/logger"
)

const (
	TICKER = "KRW-AQT"
)

func (c *Client) VAPrice() (price float32, err error) {
	logger.Debug("Try to get virtual asset price from upbit")

	ticker, err := c.call(TICKER)
	if err != nil {
		return
	}

	for i := 0; i < len(ticker); i++ {
		price = ticker[i].TradePrice

		logger.Info(ticker[i].Market + " price : " + fmt.Sprintf("%f", ticker[i].TradePrice))
	}

	return
}
