package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gravision/alphaquark-upbit-api/logger"
)

func GetResponse(req *http.Request, conn *http.Client) (buf bytes.Buffer, err error) {
	resp, err := conn.Do(req)
	if err != nil {
		logger.ErrorField("error", err.Error(), "send Request")
		return
	}

	defer resp.Body.Close()

	if _, err = io.Copy(&buf, resp.Body); err != nil {
		logger.ErrorField("error", err.Error(), "io.copy")
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", resp.StatusCode, resp.Status, buf.String())
		logger.ErrorField("error", err.Error(), "statusCode")
		return
	}

	return
}
