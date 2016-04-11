package census

import (
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Minute,
	}
)

type Client struct {
	Game      string
	ServiceID string
}

func (cl Client) buildURL(verb, col string, opts ...URLOption) string {
	return BuildURL(cl.ServiceID, "get", cl.Game, col, opts...).String()
}

func (cl Client) Get(dst interface{}, col string, opts ...URLOption) error {
	url := cl.buildURL("get", col, opts...)
	rsp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	panic("Not implemented.")
}

type CensusError struct {
	Err string `json:"error"`
}

func (err CensusError) Error() string {
	return err.Err
}
