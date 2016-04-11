package census

import (
	"encoding/json"
	"net/http"
)

// A Client is the main interface for the Census API.
type Client struct {
	// ServiceID is the ID to use when making API requests. If it is
	// blank, DefaultServiceID is used.
	ServiceID string

	// Game is the namespace to use. Examples are `ps2`, `ps2ps4us`,
	// `eq2`, etc.
	Game string

	// HTTPClient is the client to use for making API requests. If it is
	// nil, http.DefaultClient is used.
	HTTPClient *http.Client
}

// buildURL runs BuildURL with the appropriate arguments pulled from
// fields of cl.
func (cl Client) buildURL(verb, col string, opts ...URLOption) string {
	return BuildURL(cl.ServiceID, verb, cl.Game, col, opts...).String()
}

// c gets the http.Client, handling the nil case.
func (cl Client) c() *http.Client {
	if cl.HTTPClient == nil {
		return http.DefaultClient
	}

	return cl.HTTPClient
}

// Fetch returns the raw JSON that the API returns for a request. If
// the API itself has an error, such as
//
//     {"error":"No data found."}
//
// then the returned error will be an *Error.
//
// For more information about the arguments, see BuildURL.
func (cl Client) Fetch(verb, col string, opts ...URLOption) ([]byte, error) {
	url := cl.buildURL(verb, col, opts...)
	rsp, err := cl.c().Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var raw json.RawMessage
	err = json.NewDecoder(rsp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	var cerr Error
	err = json.Unmarshal(raw, &cerr)
	if err != nil {
		return nil, err
	}
	if cerr.Err != "" {
		return nil, &cerr
	}

	return raw, nil
}

// Get performs a request to the API using a `get` verb. It decodes
// the returned JSON into dst.
//
// Like with Fetch, if the API yields an error, the returned error
// will be of type *Error.
func (cl Client) Get(dst interface{}, col string, opts ...URLOption) error {
	raw, err := cl.Fetch("get", col, opts...)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, dst)
}

// Count performs a request to the API using a `count` verb. It
// returns the returned count and an error. If an error is
// encountered, the value of the returned int is undefined.
//
// Like with Fetch, if the API yields an error, the returned error
// will be of type *Error.
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

// An Error represents an error that is returned by the API.
type Error struct {
	Err string `json:"error"`
}

func (err Error) Error() string {
	return err.Err
}
