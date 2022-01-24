package etherscan

import "encoding/json"

type response struct {
	// 1 for good, 0 for error
	Status int `json:"status,string"`
	// OK for good, other words when Status equals 0
	Message string `json:"message"`
	// where response lies
	Result json.RawMessage `json:"result"`
}
