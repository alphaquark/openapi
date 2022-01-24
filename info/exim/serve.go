package exim

import "os"

func Serve() (rates map[string]float64, err error) {
	EXIM_API_KEY := os.Getenv("EXIM_API_KEY")
	exim := initClient(EXIM_API_KEY)
	rates, err = exim.ExchangeRate()
	return
}
