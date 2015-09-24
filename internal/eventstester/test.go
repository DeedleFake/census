package main

import (
	"github.com/DeedleFake/census/ps2"
	"github.com/DeedleFake/census/ps2/events"
	"log"
	"strconv"
)

func init() {
	log.SetFlags(0)
}

func main() {
	ec, err := events.NewClient("", "", "example")
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	err = ec.Subscribe(events.Sub{
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

	var c ps2.Client
	for {
		ev, err := ec.Next()
		if err != nil {
			panic(err)
		}

		switch ev := ev.(type) {
		case *events.FacilityControl:
			switch ev.NewFactionID {
			case ev.OldFactionID:
				log.Printf("%v: The %v maintainted ownership of %v on %v.\n",
					ev.WorldID,
					ev.NewFactionID,
					ev.FacilityID,
					ev.ZoneID,
				)
			default:
				log.Printf("%v: The %v captured %v on %v from the %v.\n",
					ev.WorldID,
					ev.NewFactionID,
					ev.FacilityID,
					ev.ZoneID,
					ev.OldFactionID,
				)
			}

		case *events.BattleRankUp:
			chars, err := c.Get().Character(
				map[string]string{
					"character_id": strconv.FormatInt(int64(ev.CharacterID), 10),
				},
				&ps2.Config{
					Show: []string{"name.first"},
				},
			)
			if err != nil {
				log.Printf("Error getting characters: %v", err)
				continue
			}

			log.Printf("Congratulations to %v of %v on reaching level %v.\n", chars[0].Name.First, ev.WorldID, ev.BattleRank)

		default:
			log.Printf("%#v\n", ev)
		}
	}
}
