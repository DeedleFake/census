package events

import (
	"encoding/json"
	"github.com/DeedleFake/census/ps2"
)

// Event represents a message read from an event stream.
type Event interface {
	Type() string
}

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
		ev = new(ConnectionStateChangedEvent)
	case "serviceStateChanged":
		ev = new(ServiceStateChangedEvent)
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
		Online HeartbeatEvent `json:"online"`
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
		ev = new(AchievementEarnedEvent)
	case "BattleRankUp":
		ev = new(BattleRankUpEvent)
	case "Death":
		ev = new(DeathEvent)
	case "FacilityControl":
		ev = new(FacilityControlEvent)
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

type ConnectionStateChangedEvent struct {
	Connected bool `json:"connected,string"`
}

func (ev ConnectionStateChangedEvent) Type() string {
	return "connectionStateChanged"
}

type ServiceStateChangedEvent struct {
	Detail string `json:"detail"`
	Online bool   `json:"online,string"`
}

func (ev ServiceStateChangedEvent) Type() string {
	return "serviceStateChanged"
}

type HeartbeatEvent struct {
	Briggs  bool `json:"EventServerEndpoint_Briggs_25,string"`
	Cobalt  bool `json:"EventServerEndpoint_Cobalt_13,string"`
	Connery bool `json:"EventServerEndpoint_Connery_1,string"`
	Emerald bool `json:"EventServerEndpoint_Emerald_17,string"`
	Jaeger  bool `json:"EventServerEndpoint_Jaeger_19,string"`
	Miller  bool `json:"EventServerEndpoint_Miller_10,string"`
}

func (ev HeartbeatEvent) Type() string {
	return "heartbeat"
}

type AchievementEarnedEvent struct {
	CharacterID   int       `json:"character_id,string"`
	Timestamp     int64     `json:"timestamp,string"`
	WorldID       ps2.World `json:"world_id,string"`
	AchievementID int       `json:"achievement_id,string"`
	ZoneID        int       `json:"zone_id,string"`
}

func (ev AchievementEarnedEvent) Type() string {
	return "AchievementEarned"
}

type BattleRankUpEvent struct {
	BattleRank  int       `json:"battle_rank,string"`
	CharacterID int       `json:"character_id,string"`
	Timestamp   int64     `json:"timestamp,string"`
	WorldID     ps2.World `json:"world_id,string"`
	ZoneID      int       `json:"zone_id,string"`
}

func (ev BattleRankUpEvent) Type() string {
	return "BattleRankUp"
}

type DeathEvent struct {
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
	ZoneID              int       `json:"zone_id,string"`
}

func (ev DeathEvent) Type() string {
	return "Death"
}

type FacilityControlEvent struct {
	DurationHeld int         `json:"duration_held,string"`
	FacilityID   int         `json:"facility_id,string"`
	NewFactionID ps2.Faction `json:"new_faction_id,string"`
	OldFactionID ps2.Faction `json:"old_faction_id,string"`
	OutfitID     int         `json:"outfit_id,string"`
	Timestamp    int64       `json:"timestamp,string"`
	WorldID      ps2.World   `json:"world_id,string"`
	ZoneID       int         `json:"zone_id,string"`
}

func (ev FacilityControlEvent) Type() string {
	return "FacilityControl"
}
