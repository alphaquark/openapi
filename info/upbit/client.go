package upbit

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gravision/alphaquark-upbit-api/logger"
	"github.com/gravision/alphaquark-upbit-api/request"
)

type Client struct {
	conn    *http.Client
	baseURL string
}

const (
	UPBIT_TICKER_URL = "https://api.upbit.com/v1/ticker?"

	TimeoutSecond = 30
)

func initClient() *Client {
	return &Client{
		conn: &http.Client{
			Timeout: TimeoutSecond * time.Second,
		},
		baseURL: UPBIT_TICKER_URL,
	}
}

func (c *Client) call(ticker string) (out []Response, err error) {
	req, err := http.NewRequest(http.MethodGet,
		c.craftURL(ticker),
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

	err = json.Unmarshal(content.Bytes(), &out)
	if err != nil {
		return
	}

	return
}

func (c *Client) craftURL(ticker string) (URL string) {
	q := url.Values{
		"markets": []string{ticker},
	}

	URL = c.baseURL + q.Encode()
	logger.DebugField("url", URL, "craft full request url")
	return
}
