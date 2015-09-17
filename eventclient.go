package census

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"net"
)

const (
	DefaultEventBase = "wss://push.planetside2.com/streaming"
	DefaultEventEnv  = "ps2"
)

type EventClient struct {
	c *websocket.Conn

	e *json.Encoder
	d *json.Decoder
}

func NewEventClient(base, env, svcid string) (*EventClient, error) {
	if base == "" {
		base = DefaultEventBase
	}
	if env == "" {
		env = DefaultEventEnv
	}

	c, err := websocket.Dial(base+"?environment="+env+"&service-id=s:"+svcid, "", "")
	if err != nil {
		return nil, err
	}

	return &EventClient{
		c: c,

		e: json.NewEncoder(c),
		d: json.NewDecoder(c),
	}, nil
}

func (c *EventClient) Conn() net.Conn {
	return c.c
}

func (c *EventClient) sub(action string, events []Sub) error {
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

func (c *EventClient) Subscribe(events ...Sub) error {
	return c.sub("subscribe", events)
}

func (c *EventClient) Unsubscribe(events ...Sub) error {
	return c.sub("clearSubscribe", events)
}
