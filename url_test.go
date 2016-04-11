package census_test

import (
	"."
	"testing"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		id, verb, game, col string
		opts                []census.URLOption
		out                 string
	}{
		{
			"id", "verb", "game", "col",
			[]census.URLOption{
				census.SearchOption("name", "val"),
			},
			"https://census.daybreakgames.com/s:id/verb/game/col?name=val",
		},
	}

	for _, test := range tests {
		url := census.BuildURL(
			test.id,
			test.verb,
			test.game,
			test.col,
			test.opts...,
		).String()

		if url != test.out {
			t.Errorf("Failed: %#v\n\tGot %q", test, url)
		}
	}
}
