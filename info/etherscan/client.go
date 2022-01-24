package etherscan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gravision/alphaquark-upbit-api/logger"
	"github.com/gravision/alphaquark-upbit-api/request"
	myutil "github.com/gravision/alphaquark-upbit-api/utils"
)

type Client struct {
	conn    *http.Client
	apiKey  string
	baseURL string
}

const (
	TimeoutSecond = 30
)

func initClient(network Network, APIKey string) *Client {
	var tBaseURL string

	subDomain := network.SubDomain()
	if subDomain == "bsc" {
		logger.Info("Selected Binance smart chain")
		tBaseURL = "https://api.bscscan.io/api?"
	} else if subDomain == "api" {
		logger.Info("Ethereum mainnet")
		tBaseURL = fmt.Sprintf("https://%s.etherscan.io/api?", network.SubDomain())
	} else {
		logger.Panic("Not implementation")
		return nil
	}

	logger.InfoField("BaseURL", tBaseURL, "Etherscan base URL")

	return &Client{
		conn: &http.Client{
			Timeout: TimeoutSecond * time.Second,
		},
		apiKey:  APIKey,
		baseURL: tBaseURL,
	}
}

// call does almost all the work.
func (c *Client) call(module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	req, err := http.NewRequest(http.MethodGet,
		c.craftURL(module, action, param),
		http.NoBody,
	)

	if err != nil {
		logger.ErrorField("error", err.Error(), "http.NewRequest")
		return
	}

	logger.Debug("try conn.Do")
	// resp, err := c.conn.Do(req)
	// if err != nil {
	// 	logger.ErrorField("error", err.Error(), "send Request")
	// 	return
	// }
	// defer resp.Body.Close()
	// var content bytes.Buffer
	// if _, err = io.Copy(&content, resp.Body); err != nil {
	// 	logger.ErrorField("error", err.Error(), "io.copy")
	// 	return
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	err = fmt.Errorf("response status %v %s, response body: %s", resp.StatusCode, resp.Status, content.String())
	// 	logger.ErrorField("error", err.Error(), "statusCode")
	// 	return
	// }
	content, err := request.GetResponse(req, c.conn)
	if err != nil {
		return
	}

	var tResponse response
	err = json.Unmarshal(content.Bytes(), &tResponse)
	if err != nil {
		return
	}
	if tResponse.Status != 1 {
		logger.ErrorField("error", err.Error(), "etherscan server")
		return
	}

	err = json.Unmarshal(tResponse.Result, outcome)
	if err != nil {
		logger.ErrorField("error", err.Error(), "unmarshal outcome")
		return
	}

	return
}

// craftURL returns desired URL via param provided
func (c *Client) craftURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{c.apiKey},
	}

	for k, v := range param {
		q[k] = myutil.ExtractValue(v)
	}

	URL = c.baseURL + q.Encode()
	logger.DebugField("url", URL, "craft full request url")
	return
}
