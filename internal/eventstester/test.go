package main

import (
	"fmt"
	"github.com/DeedleFake/census/ps2/events"
)

func main() {
	c, err := events.NewClient("", "", "example")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	err = c.Subscribe(events.Sub{
		Events: []string{
			"BattleRankUp",
			"FacilityControl",
		},
		Chars:  events.SubAll,
		Worlds: events.SubAll,
	})
	if err != nil {
		panic(err)
	}

	for {
		ev, err := c.Next()
		if err != nil {
			panic(err)
		}

		switch ev := ev.(type) {
		case *events.FacilityControlEvent:
			switch ev.NewFactionID {
			case ev.OldFactionID:
				fmt.Printf("%v: The %v maintainted ownership of %v on %v.\n",
					ev.WorldID,
					ev.NewFactionID,
					ev.FacilityID,
					ev.ZoneID,
				)
			default:
				fmt.Printf("%v: The %v captured %v on %v from the %v.\n",
					ev.WorldID,
					ev.NewFactionID,
					ev.FacilityID,
					ev.ZoneID,
					ev.OldFactionID,
				)
			}

		case *events.BattleRankUpEvent:
			if ev.BattleRank >= 90 {
				fmt.Printf("Congratulations to %v on reaching level %v.\n", ev.CharacterID, ev.BattleRank)
			}

		default:
			fmt.Printf("%#v\n", ev)
		}
	}
}
