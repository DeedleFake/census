package ps2

import (
	"strconv"
)

// World is a convenience type for dealing with PlanetSide 2 world IDs.
type World int64

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

// Faction is a convenience type for dealing with PlanetSide 2 faction
// IDs.
type Faction int64

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

// Zone is a convenience type for dealing PlanetSide 2 zone IDs.
type Zone int64

const (
	Indar        Zone = 2
	Hossin       Zone = 4
	Amerish      Zone = 6
	Esamir       Zone = 8
	VRTrainingNC Zone = 96
	VRTrainingTR Zone = 97
	VRTrainingVS Zone = 98
	Cleanroom    Zone = 200
)

func (z Zone) String() string {
	switch z {
	case Indar:
		return "Indar"
	case Hossin:
		return "Hossin"
	case Amerish:
		return "Amerish"
	case Esamir:
		return "Esamir"
	case VRTrainingNC:
		return "VR Training (NC)"
	case VRTrainingTR:
		return "VR Training (TR)"
	case VRTrainingVS:
		return "VR Training (VS)"
	case Cleanroom:
		return "Cleanroom"
	}

	panic("Unkown zone ID: " + strconv.FormatInt(int64(z), 10))
}
