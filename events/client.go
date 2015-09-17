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

// Client is a connection to the Census API event streaming system.
type Client struct {
	c *websocket.Conn

	e *json.Encoder
	d *json.Decoder
}

// NewClient creates a new Client. It connects to the URL at base and
// the environment env using the service ID svcid. If base or env are
// the empty string, then DefaultBase and DefaultEnv are used,
// respectively.
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

// Close closes the connection to the server. Use of the Client after
// calling Close may or may not panic.
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

// Conn returns the underlying connection to the server.
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

// Sub represents a subscription request.
type Sub struct {
	Events []string
	Chars  []string
	Worlds []string
}

var (
	// SubAll is a special value that can be used for any or all of the
	// fields of a Sub as a means of matching all possible values for
	// that field. For example,
	//
	//     Sub{
	//         Events: []string{"BattleRankUp"},
	//         Chars: SubAll,
	//     }
	//
	// will subscribe to level up events for every character.
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

// Subscribe subscribes to the events specified by sub. For more
// information, see http://census.daybreakgames.com/#websocket-details
func (c *Client) Subscribe(sub Sub) error {
	return c.sub("subscribe", sub.Events, sub.Chars, sub.Worlds)
}

// Unsubscribe unsubscribes from the events specified by sub. For more
// information, see http://census.daybreakgames.com/#websocket-details
func (c *Client) Unsubscribe(sub Sub) error {
	return c.sub("clearSubscribe", sub.Events, sub.Chars, sub.Worlds)
}

// Next blocks until the next event can be read from the event stream
// or an error occurs. It then returns that event or the error. If an
// unsupported event is read, a nil Event and an UnknownEventTypeError
// are returned.
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
