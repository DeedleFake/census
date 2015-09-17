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

	panic("Unknown world ID: " + strconv.FormatInt(int64(w), 10))
}

type PS2Faction int

const (
	NaniteSystems   PS2Faction = 0
	VanuSovereignty PS2Faction = 1
	NewConglomerate PS2Faction = 2
	TerranRepublic  PS2Faction = 3
)

func (f PS2Faction) String() string {
	switch f {
	case NaniteSystems:
		return "Nanite Systems"
	case VanuSovereignty:
		return "Vanu Sovereignty"
	case NewConglomerate:
		return "New Conglomerate"
	case TerranRepublic:
		return "Terran Republic"
	}

	panic("Unknown faction ID: " + strconv.FormatInt(int64(f), 10))
}
