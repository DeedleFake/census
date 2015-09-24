package ps2

import (
	"encoding/json"
	"io/ioutil"
)

type Get struct {
	c *Client
}

type Character struct {
	CharacterID int64 `json:"character_id,string"`
	Name        struct {
		First      string `json:"first"`
		FirstLower string `json:"first_lower"`
	} `json:"name"`
	FactionID Faction `json:"faction_id,string"`
	HeadID    int     `json:"head_id,string"`
	TitleID   int     `json:"title_id,string"`
	Times     struct {
		Creation      int64 `json:"creation,string"`
		LastSave      int64 `json:"last_save,string"`
		LastLogin     int64 `json:"last_login,string"`
		LoginCount    int   `json:"login_count,string"`
		MinutesPlayed int   `json:"minutes_played,string"`
	} `json:"times"`
	Certs struct {
		EarnedPoints    int     `json:"earned_points,string"`
		GiftedPoints    int     `json:"gifted_points,string"`
		SpentPoints     int     `json:"spent_points,string"`
		AvailablePoints int     `json:"available_points,string"`
		PercentToNext   float64 `json:"percent_to_next,string"`
	} `json:"certs"`
	BattleRank struct {
		PercentToNext float64 `json:"percent_to_next,string"`
		Value         int     `json:"value,string"`
	} `json:"battle_rank"`
	ProfileID   int `json:"profile_id,string"`
	DailyRibbon struct {
		Count int `json:"count,string"`
	} `json:"daily_ribbon"`
}

func (g *Get) Character(search map[string]string, config *Config) ([]Character, error) {
	rsp, err := g.c.c().Get(g.c.buildURL("get", "character", search, config))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// Read into buffer so that it can be reread. This will be used
	// later for resolves.
	raw, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		List []Character `json:"character_list"`
	}
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	return data.List, nil
}
