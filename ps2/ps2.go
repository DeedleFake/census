package ps2

import (
	"strconv"
)

// World is a convience type for dealing with PlanetSide 2 world IDs.
type World int

const (
	Connery World = 1
	Miller  World = 10
	Cobalt  World = 13
	Emerald World = 17
	Jaeger  World = 19
	Briggs  World = 25
)

func (w World) String() string {
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

// Faction is a convience type for dealing with PlanetSide 2 faction
// IDs.
type Faction int

const (
	NaniteSystems   Faction = 0
	VanuSovereignty Faction = 1
	NewConglomerate Faction = 2
	TerranRepublic  Faction = 3
)

func (f Faction) String() string {
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
