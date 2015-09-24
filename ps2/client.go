package ps2

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func (c *Client) buildURL(t string, col string, search map[string]string, config *Config) string {
	q := make(url.Values)

	for k, v := range search {
		q.Add(k, v)
	}

	config.addToQuery(q)

	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()

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
	Show            []string
	Hide            []string
	Sort            Sort
	Has             []string
	IgnoreCase      bool
	Limit           int
	LimitPerDB      int
	Start           int
	Lang            string
	ExactMatchFirst bool
	TryOnce         bool
}

func (c *Config) addToQuery(q url.Values) {
	if c == nil {
		return
	}

	if len(c.Show) > 0 {
		q.Set("c:show", strings.Join(c.Show, ","))
	}

	if len(c.Hide) > 0 {
		q.Set("c:hide", strings.Join(c.Hide, ","))
	}

	c.Sort.addToQuery(q)

	if len(c.Has) > 0 {
		q.Set("c:has", strings.Join(c.Has, ","))
	}

	if c.IgnoreCase {
		q.Set("c:case", "false")
	}

	if c.Limit > 0 {
		q.Set("c:limit", strconv.FormatInt(int64(c.Limit), 10))
	}

	if c.LimitPerDB > 0 {
		q.Set("c:limitPerDB", strconv.FormatInt(int64(c.LimitPerDB), 10))
	}

	if c.Start > 0 {
		q.Set("c:start", strconv.FormatInt(int64(c.Start), 10))
	}

	if c.Lang != "" {
		q.Set("c:lang", c.Lang)
	}

	if c.ExactMatchFirst {
		q.Set("c:exactMatchFirst", "true")
	}

	if c.TryOnce {
		q.Set("c:retry", "false")
	}
}

type Sort []SortDir

func (s Sort) addToQuery(q url.Values) {
	if len(s) == 0 {
		return
	}

	param := make([]string, 0, len(s))
	for _, dir := range s {
		param = append(param, dir.String())
	}

	q.Set("c:sort", strings.Join(param, ","))
}

type SortDir struct {
	Field string
	Dir   int
}

func (dir SortDir) String() string {
	d := "1"
	if dir.Dir < 0 {
		d = "-1"
	}

	return dir.Field + ":" + d
}
