package events

import (
	"encoding/json"
	"github.com/DeedleFake/census/ps2"
)

// Event represents a message read from an event stream.
type Event interface{}

func eventFromRaw(raw []byte) (Event, error) {
	var common struct {
		Service string `json:"service"`
		Type    string `json:"type"`
	}
	err := json.Unmarshal(raw, &common)
	if err != nil {
		return nil, err
	}

	var ev Event
	switch common.Type {
	case "connectionStateChanged":
		ev = new(ConnectionStateChanged)
	case "serviceStateChanged":
		ev = new(ServiceStateChanged)
	case "heartbeat":
		return heartbeat(raw)
	case "serviceMessage":
		return serviceMessage(raw)
	case "":
		return nil, nil
	default:
		return nil, UnknownEventTypeError(common.Type)
	}

	err = json.Unmarshal(raw, &ev)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func heartbeat(raw []byte) (Event, error) {
	var outer struct {
		Online Heartbeat `json:"online"`
	}
	err := json.Unmarshal(raw, &outer)
	if err != nil {
		return nil, err
	}

	return outer.Online, nil
}

func serviceMessage(raw []byte) (Event, error) {
	var outer struct {
		Payload json.RawMessage `json:"payload"`
	}
	err := json.Unmarshal(raw, &outer)
	if err != nil {
		return nil, err
	}

	var common struct {
		EventName string `json:"event_name"`
	}
	err = json.Unmarshal(outer.Payload, &common)
	if err != nil {
		return nil, err
	}

	var ev Event
	switch common.EventName {
	case "AchievementEarned":
		ev = new(AchievementEarned)
	case "BattleRankUp":
		ev = new(BattleRankUp)
	case "Death":
		ev = new(Death)
	case "FacilityControl":
		ev = new(FacilityControl)
	case "GainExperience":
		ev = new(GainExperience)
	default:
		return nil, UnknownEventTypeError(common.EventName)
	}

	err = json.Unmarshal(outer.Payload, ev)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

// UnknownEventTypeError is returned by event stream readers that
// encountered an event that they didn't know how to handle. The value
// of the error is set to the string name of the event type.
type UnknownEventTypeError string

func (err UnknownEventTypeError) Error() string {
	return "Unknown event type: " + string(err)
}

type ConnectionStateChanged struct {
	Connected bool `json:"connected,string"`
}

type ServiceStateChanged struct {
	Detail string `json:"detail"`
	Online bool   `json:"online,string"`
}

type Heartbeat struct {
	Briggs  bool `json:"EventServerEndpoint_Briggs_25,string"`
	Cobalt  bool `json:"EventServerEndpoint_Cobalt_13,string"`
	Connery bool `json:"EventServerEndpoint_Connery_1,string"`
	Emerald bool `json:"EventServerEndpoint_Emerald_17,string"`
	Jaeger  bool `json:"EventServerEndpoint_Jaeger_19,string"`
	Miller  bool `json:"EventServerEndpoint_Miller_10,string"`
}

type AchievementEarned struct {
	CharacterID   int       `json:"character_id,string"`
	Timestamp     int64     `json:"timestamp,string"`
	WorldID       ps2.World `json:"world_id,string"`
	AchievementID int       `json:"achievement_id,string"`
	ZoneID        ps2.Zone  `json:"zone_id,string"`
}

type BattleRankUp struct {
	BattleRank  int       `json:"battle_rank,string"`
	CharacterID int       `json:"character_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
	ZoneID      ps2.Zone  `json:"zone_id,string"`
}

type Death struct {
	AttackerCharacterID int       `json:"attacker_character_id,string"`
	AttackerFireModeID  int       `json:"attacker_fire_mode_id,string"`
	AttackerLoadoutID   int       `json:"attacker_loadout_id,string"`
	AttackerVehicleID   int       `json:"attacker_vehicle_id,string"`
	AttackerWeaponID    int       `json:"attacker_weapon_id,string"`
	CharacterID         int       `json:"character_id,string"`
	CharacterLoadoutID  int       `json:"character_loadout_id,string"`
	IsCritical          int       `json:"is_critical,string"`
	IsHeadshot          int       `json:"is_headshot,string"`
	Timestamp           int64     `json:"timestamp,string"`
	VehicleID           int       `json:"vehicle_id,string"`
	WorldID             ps2.World `json:"world_id,string"`
	ZoneID              ps2.Zone  `json:"zone_id,string"`
}

type FacilityControl struct {
	DurationHeld int         `json:"duration_held,string"`
	FacilityID   int         `json:"facility_id,string"`
	NewFactionID ps2.Faction `json:"new_faction_id,string"`
	OldFactionID ps2.Faction `json:"old_faction_id,string"`
	OutfitID     int         `json:"outfit_id,string"`
	Timestamp    int64       `json:"timestamp,string"`
	WorldID      ps2.World   `json:"world_id,string"`
	ZoneID       ps2.Zone    `json:"zone_id,string"`
}

type GainExperience struct {
	Amount       int       `json:"amount,string"`
	CharacterID  int       `json:"character_id,string"`
	ExperienceID int       `json:"experience_id,string"`
	LoadoutID    int       `json:"loadout_id,string"`
	OtherID      int       `json:"other_id,string"`
	Timestamp    int64     `json:"timestamp,string"`
	WorldID      ps2.World `json:"world_id,string"`
	ZoneID       ps2.Zone  `json:"zone_id,string"`
}

type ItemAdded struct {
	CharacterID int      `json:"character_id,string"`
	Context     string   `json:"context"`
	ItemCount   int      `json:"item_count,string"`
	ItemID      int      `json:"item_id,string"`
	Timestamp   int64    `json:"timestamp,string"`
	WorldID     int      `json:"world_id,string"`
	ZoneID      ps2.Zone `json:"zone_id,string"`
}

// I'm not sure what these types should be.
//type MetagameEvent struct {
//	ExperienceBonus string `json:"experience_bonus"`
//	FactionNc       string `json:"faction_nc"`
//	FactionTr       string `json:"faction_tr"`
//	FactionVs       string `json:"faction_vs"`
//	MetagameID      string `json:"metagame_event_id"`
//	MetagameState   string `json:"metagame_event_state"`
//	Timestamp       string `json:"timestamp"`
//	WorldID         string `json:"world_id"`
//	ZoneID          string `json:"zone_id"`
//}

type PlayerFacilityCapture struct {
	CharacterID int       `json:"character_id,string"`
	FacilityID  int       `json:"facility_id,string"`
	OutfitID    int       `json:"outfit_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
	ZoneID      ps2.Zone  `json:"zone_id,string"`
}

type PlayerFacilityDefend struct {
	CharacterID int       `json:"character_id,string"`
	FacilityID  int       `json:"facility_id,string"`
	OutfitID    int       `json:"outfit_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
	ZoneID      ps2.Zone  `json:"zone_id,string"`
}

type PlayerLogin struct {
	CharacterID int       `json:"character_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
}

type PlayerLogout struct {
	CharacterID int       `json:"character_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
}

type SkillAdded struct {
	CharacterID int       `json:"character_id,string"`
	SkillID     int       `json:"skill_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
	ZoneID      ps2.Zone  `json:"zone_id,string"`
}

type VehicleDestroy struct {
	AttackerCharacterID int         `json:"attacker_character_id,string"`
	AttackerLoadoutID   int         `json:"attacker_loadout_id,string"`
	AttackerVehicleID   int         `json:"attacker_vehicle_id,string"`
	AttackerWeaponID    int         `json:"attacker_weapon_id,string"`
	CharacterID         int         `json:"character_id,string"`
	FacilityID          int         `json:"facility_id,string"`
	FactionID           ps2.Faction `json:"faction_id,string"`
	Timestamp           int64       `json:"timestamp,string"`
	VehicleID           int         `json:"vehicle_id,string"`
	WorldID             ps2.World   `json:"world_id,string"`
	ZoneID              ps2.Zone    `json:"zone_id,string"`
}
