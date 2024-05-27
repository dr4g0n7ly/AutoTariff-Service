package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) AggregateInvoice(distance types.Distance) error {
	httpc := http.DefaultClient
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}
	res, err := httpc.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with non 200 status code %d", res.StatusCode)
	}
	return nil
}
