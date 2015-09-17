package events

import (
	"encoding/json"
)

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

	var ev interface{}
	switch common.Type {
	case "connectionStateChanged":
		ev = new(ConnectionStateChangedEvent)
	case "serviceStateChanged":
		ev = new(ServiceStateChangedEvent)
	case "heartbeat":
		ev = new(HeartbeatEvent)
	case "":
		return nil, nil
	default:
		return nil, UnknownEventType(common.Type)
	}

	err = json.Unmarshal(raw, &ev)
	if err != nil {
		return nil, err
	}

	return ev.(Event), nil
}

type UnknownEventType string

func (err UnknownEventType) Error() string {
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

type DeathEvent struct {
	AttackerCharacterID string `json:"attacker_character_id"`
	AttackerFireModeID  string `json:"attacker_fire_mode_id"`
	AttackerLoadoutID   string `json:"attacker_loadout_id"`
	AttackerVehicleID   string `json:"attacker_vehicle_id"`
	AttackerWeaponID    string `json:"attacker_weapon_id"`
	CharacterID         string `json:"character_id"`
	CharacterLoadoutID  string `json:"character_loadout_id"`
	EventName           string `json:"event_name"`
	IsCritical          bool   `json:"is_critical,string"`
	IsHeadshot          bool   `json:"is_headshot,string"`
	Timestamp           string `json:"timestamp"`
	VehicleID           string `json:"vehicle_id"`
	WorldID             string `json:"world_id"`
	ZoneID              string `json:"zone_id"`
}
