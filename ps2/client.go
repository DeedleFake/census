package ps2

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
)

const (
	DefaultBase      = "http://census.daybreakgames.com"
	DefaultDomain    = "ps2"
	DefaultServiceID = "example"
)

type Client struct {
	C         *http.Client
	Base      string
	ServiceID string
	Domain    string
}

func (c *Client) c() *http.Client {
	if c.C == nil {
		return http.DefaultClient
	}

	return c.C
}

func (c *Client) base() string {
	if c.Base == "" {
		return DefaultBase
	}

	return c.Base
}

func (c *Client) domain() string {
	if c.Domain == "" {
		return DefaultDomain
	}

	return c.Domain
}

func (c *Client) serviceID() string {
	if c.ServiceID == "" {
		return DefaultServiceID
	}

	return c.ServiceID
}

func (c *Client) buildURL(t string, col string, search map[string]string, config *Config) string {
	q := make(url.Values)

	for k, v := range search {
		q.Add(k, v)
	}

	config.addToQuery(q)

	var buf bytes.Buffer
	buf.WriteString(c.base())
	buf.WriteString("/s:")
	buf.WriteString(c.serviceID())
	buf.WriteByte('/')
	buf.WriteString(t)
	buf.WriteByte('/')
	buf.WriteString(c.domain())
	buf.WriteByte('/')
	buf.WriteString(col)
	buf.WriteByte('?')
	buf.WriteString(q.Encode())

	return buf.String()
}

func (c *Client) Get() *Get {
	return &Get{
		c: c,
	}
}

type Config struct {
	Show []string
}

func (c *Config) addToQuery(q url.Values) {
	if c == nil {
		return
	}

	if len(c.Show) > 0 {
		q.Add("c:show", strings.Join(c.Show, ","))
	}
}
