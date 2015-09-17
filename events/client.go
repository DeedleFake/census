package events

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"net"
)

const (
	DefaultBase = "wss://push.planetside2.com/streaming"
	DefaultEnv  = "ps2"
)

type Client struct {
	c *websocket.Conn

	e *json.Encoder
	d *json.Decoder
}

func NewClient(base, env, svcid string) (*Client, error) {
	if base == "" {
		base = DefaultBase
	}
	if env == "" {
		env = DefaultEnv
	}

	c, err := websocket.Dial(base+"?environment="+env+"&service-id=s:"+svcid, "", "http://localhost")
	if err != nil {
		return nil, err
	}

	return &Client{
		c: c,

		e: json.NewEncoder(c),
		d: json.NewDecoder(c),
	}, nil
}

func (c *Client) Close() error {
	err := c.c.Close()
	if err != nil {
		return err
	}

	c.c = nil
	c.e = nil
	c.d = nil

	return nil
}

func (c *Client) Conn() net.Conn {
	return c.c
}

// This is somewhat problematic as it causes arbitrary JSON to get
// sent as part of the event stream.
func (c *Client) echo(payload interface{}) error {
	return c.e.Encode(&struct {
		Service string      `json:"service"`
		Action  string      `json:"action"`
		Payload interface{} `json:"payload"`
	}{
		"event",
		"echo",
		payload,
	})
}

type Sub struct {
	Events []string
	Chars  []string
	Worlds []string
}

var (
	SubAll = []string{"all"}
)

func (c *Client) sub(action string, events, chars, worlds []string) error {
	return c.e.Encode(&struct {
		Service string   `json:"service"`
		Action  string   `json:"action"`
		Events  []string `json:"eventNames,omitempty"`
		Chars   []string `json:"characters,omitempty"`
		Worlds  []string `json:"worlds,omitempty"`
	}{
		"event",
		action,
		events,
		chars,
		worlds,
	})
}

func (c *Client) Subscribe(sub Sub) error {
	return c.sub("subscribe", sub.Events, sub.Chars, sub.Worlds)
}

func (c *Client) Unsubscribe(sub Sub) error {
	return c.sub("clearSubscribe", sub.Events, sub.Chars, sub.Worlds)
}

func (c *Client) Next() (ev Event, err error) {
	for (ev == nil) && (err == nil) {
		var raw json.RawMessage
		err = c.d.Decode(&raw)
		if err != nil {
			return
		}

		ev, err = eventFromRaw(raw)
	}

	return
}
