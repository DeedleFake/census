package census

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Minute,
	}
)

type Client struct {
	ServiceID string
	Game      string
}

func (cl Client) buildURL(verb, col string, opts ...URLOption) string {
	return BuildURL(cl.ServiceID, verb, cl.Game, col, opts...).String()
}

func (cl Client) Fetch(verb, col string, opts ...URLOption) (json.RawMessage, error) {
	url := cl.buildURL(verb, col, opts...)
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var raw json.RawMessage
	err = json.NewDecoder(rsp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	var cerr CensusError
	err = json.Unmarshal(raw, &cerr)
	if err != nil {
		return nil, err
	}
	if cerr.Err != "" {
		return nil, &cerr
	}

	return raw, nil
}

func (cl Client) Get(dst interface{}, col string, opts ...URLOption) error {
	raw, err := cl.Fetch("get", col, opts...)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, dst)
}

func (cl Client) Count(col string, opts ...URLOption) (int, error) {
	raw, err := cl.Fetch("count", col, opts...)
	if err != nil {
		return 0, err
	}

	var data struct {
		Count int `json:"count"`
	}
	err = json.Unmarshal(raw, &data)
	return data.Count, err
}

type CensusError struct {
	Err string `json:"error"`
}

func (err CensusError) Error() string {
	return err.Err
}
