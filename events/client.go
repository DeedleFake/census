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

func (c *Client) sub(action string, events []Sub) error {
	if len(events) == 0 {
		return nil
	}

	sub := map[string]interface{}{
		"service": "event",
		"action":  action,
	}

	ev := make([]string, len(events))
	for _, e := range events {
		ev = append(ev, e.name())

		for pname, pslice := range e.params() {
			var s []string
			if i, ok := sub[pname]; ok {
				s = i.([]string)
			}

			sub[pname] = append(s, pslice...)
		}
	}
	sub["eventNames"] = ev

	return c.e.Encode(sub)
}

func (c *Client) Subscribe(events ...Sub) error {
	return c.sub("subscribe", events)
}

func (c *Client) Unsubscribe(events ...Sub) error {
	return c.sub("clearSubscribe", events)
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
