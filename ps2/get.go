package ps2

import (
	"encoding/json"
)

type Get struct {
	c *Client
}

func (g *Get) Custom(out interface{}, col string, search map[string]string, config *Config) error {
	rsp, err := g.c.c().Get(g.c.buildURL("get", col, search, config))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	dec := json.NewDecoder(rsp.Body)

	var data map[string]json.RawMessage
	err = dec.Decode(&data)
	if err != nil {
		return err
	}

	list, ok := data[col+"_list"]
	if !ok {
		return UnknownCollectionError(col)
	}

	return json.Unmarshal(list, out)
}
