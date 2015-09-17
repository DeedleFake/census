package census

type Event interface {
}

type rawEvent struct {
	Service string `json:"service"`
	Type    string `json:"type"`
}

type ConnectionStateChangedEvent struct {
}

type HeartbeatEvent struct {
	Briggs  bool `json:"EventServerEndpoint_Briggs_25,string"`
	Cobalt  bool `json:"EventServerEndpoint_Cobalt_13,string"`
	Connery bool `json:"EventServerEndpoint_Connery_1,string"`
	Emerald bool `json:"EventServerEndpoint_Emerald_17,string"`
	Jaeger  bool `json:"EventServerEndpoint_Jaeger_19,string"`
	Miller  bool `json:"EventServerEndpoint_Miller_10,string"`
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
