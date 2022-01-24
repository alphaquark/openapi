package exim

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/gravision/alphaquark-upbit-api/logger"
	"github.com/gravision/alphaquark-upbit-api/request"
)

type Client struct {
	conn    *http.Client
	apiKey  string
	baseURL string
}

var (
	exim_apiKey string
)

const (
	TimeoutSecond = 30
)

func initClient(APIKey string) *Client {
	exim_apiKey = APIKey
	return &Client{
		conn: &http.Client{
			Timeout: TimeoutSecond * time.Second,
		},
		apiKey:  APIKey,
		baseURL: "https://www.koreaexim.go.kr/site/program/financial/exchangeJSON?",
	}
}

func (c *Client) call(authKey, data string) (rates []Response, err error) {
	req, err := http.NewRequest(http.MethodGet,
		c.craftURL(authKey, data),
		nil,
	)

	if err != nil {
		logger.ErrorField("error", err.Error(), "http.NewRequest")
		return
	}

	content, err := request.GetResponse(req, c.conn)
	if err != nil {
		return
	}

	err = json.Unmarshal(content.Bytes(), &rates)
	if err != nil {
		return
	}

	if len(rates) == 0 {
		err = errors.New("exim server is not ready yet")
	} else {
		if rates[0].RESULT != 1 {
			logger.ErrorField("error", err.Error(), "exim server")
			return
		}
	}

	return
}

func (c *Client) craftURL(authKey, data string) (URL string) {
	q := url.Values{
		"authkey": []string{authKey},
		"data":    []string{data},
	}

	URL = c.baseURL + q.Encode()
	logger.DebugField("url", URL, "craft full request url")
	return
}
