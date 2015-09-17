package main

import (
	"fmt"
	"github.com/DeedleFake/census/events"
)

func main() {
	c, err := events.NewClient("", "", "example")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	err = c.Subscribe(&events.WorldSub{
		Event:  "FacilityControl",
		Worlds: []string{"all"},
	})
	if err != nil {
		panic(err)
	}

	for {
		ev, err := c.Next()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n", ev)
	}
}
