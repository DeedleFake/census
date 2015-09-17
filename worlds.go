package census

import (
	"strconv"
)

// PS2World is a convience type for dealing with PlanetSide 2 world
// IDs.
type PS2World int

const (
	Connery PS2World = 1
	Miller  PS2World = 10
	Cobalt  PS2World = 13
	Emerald PS2World = 17
	Jaeger  PS2World = 19
	Briggs  PS2World = 25
)

func (w PS2World) String() string {
	switch w {
	case Connery:
		return "Connery"
	case Miller:
		return "Miller"
	case Cobalt:
		return "Cobalt"
	case Emerald:
		return "Emerald"
	case Jaeger:
		return "Jaeger"
	case Briggs:
		return "Briggs"
	}

	panic("Unknown world: " + strconv.FormatInt(int64(w), 10))
}
