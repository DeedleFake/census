package ps2

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

//func TestBuildURL(t *testing.T) {
//	var c Client
//	built := c.buildURL(
//		"get",
//		"character",
//		map[string]string{
//			"name.first": "DeedleFakeTR",
//		},
//		&Config{
//			Show: []string{"name", "battle_rank"},
//		},
//	)
//	ex := "http://census.daybreakgames.com/s:example/get/ps2/character?name.first=DeedleFakeTR&c:show=name,battle_rank"
//
//	if built != ex {
//		t.Errorf("Got %q", built)
//		t.Errorf("Expected %q", ex)
//	}
//}

func TestConfigAddToQuery(t *testing.T) {
	c := &Config{
		Show: []string{"show1", "show2"},
		Hide: []string{"hide1", "hide2"},
		Sort: Sort{
			{
				Field: "sort1",
			},
			{
				Field: "sort2",
				Dir:   -1,
			},
		},
		Has:             []string{"has1", "has2"},
		Resolve:         []string{"resolve1", "resolve2"},
		IgnoreCase:      true,
		Limit:           3,
		LimitPerDB:      9,
		Start:           30,
		Lang:            "en",
		ExactMatchFirst: true,
		TryOnce:         true,
	}

	got := make(url.Values, 11)
	c.addToQuery(got)

	ex := make(url.Values, 11)
	ex.Set("c:show", "show1,show2")
	ex.Set("c:hide", "hide1,hide2")
	ex.Set("c:sort", "sort1:1,sort2:-1")
	ex.Set("c:has", "has1,has2")
	ex.Set("c:resolve", "resolve1,resolve2")
	ex.Set("c:case", "false")
	ex.Set("c:limit", "3")
	ex.Set("c:limitPerDB", "9")
	ex.Set("c:start", "30")
	ex.Set("c:lang", "en")
	ex.Set("c:exactMatchFirst", "true")
	ex.Set("c:retry", "false")

	if !reflect.DeepEqual(got, ex) {
		t.Errorf("Got %#v", got)
		t.Errorf("Expected: %#v", ex)
	}
}

func ExampleClient() {
	var c Client
	chars, err := c.Get().Character(
		map[string]string{
			"name.first": "^DeedleFake",
		},
		&Config{
			Sort: Sort{
				{
					Field: "times.creation",
				},
			},
			Resolve: []string{
				"faction",
			},
			Limit: 3,
		},
	)
	if err != nil {
		panic(err)
	}

	for _, c := range chars {
		fmt.Printf("%v (%v)\n", c.Name.First, c.Faction.CodeTag)
	}

	// Output: DeedleFake (NC)
	// DeedleFakeConnery (VS)
	// DeedleFakeTR (TR)
}
