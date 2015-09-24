package ps2

import (
	"fmt"
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

func ExampleClient() {
	var c Client
	chars, err := c.Get().Character(
		map[string]string{
			"name.first": "DeedleFakeTR",
		},
		&Config{
			Show: []string{
				"name",
				"battle_rank",
			},
		},
	)
	if err != nil {
		panic(err)
	}

	for _, c := range chars {
		fmt.Printf("%v (%v)\n", c.Name.First, c.BattleRank.Value)
	}

	// Output: DeedleFakeTR (100)
}
