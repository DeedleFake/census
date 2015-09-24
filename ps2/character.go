package ps2

type Character struct {
	CharacterID int64 `json:"character_id,string"`
	Name        struct {
		First      string `json:"first"`
		FirstLower string `json:"first_lower"`
	} `json:"name"`
	FactionID Faction `json:"faction_id,string"`
	HeadID    int     `json:"head_id,string"`
	TitleID   int     `json:"title_id,string"`
	Times     struct {
		Creation      int64 `json:"creation,string"`
		LastSave      int64 `json:"last_save,string"`
		LastLogin     int64 `json:"last_login,string"`
		LoginCount    int   `json:"login_count,string"`
		MinutesPlayed int   `json:"minutes_played,string"`
	} `json:"times"`
	Certs struct {
		EarnedPoints    int     `json:"earned_points,string"`
		GiftedPoints    int     `json:"gifted_points,string"`
		SpentPoints     int     `json:"spent_points,string"`
		AvailablePoints int     `json:"available_points,string"`
		PercentToNext   float64 `json:"percent_to_next,string"`
	} `json:"certs"`
	BattleRank struct {
		PercentToNext float64 `json:"percent_to_next,string"`
		Value         int     `json:"value,string"`
	} `json:"battle_rank"`
	ProfileID   int `json:"profile_id,string"`
	DailyRibbon struct {
		Count int `json:"count,string"`
	} `json:"daily_ribbon"`

	Items []struct {
		ItemID     int64 `json:"item_id,string"`
		StackCount int   `json:"stack_count,string"`
	} `json:"items"`

	ItemsFull []struct {
		ItemID              int64             `json:"item_id,string"`
		StackCount          int               `json:"stack_count,string"`
		ItemTypeID          int64             `json:"item_type_id,string"`
		ItemCategoryID      int64             `json:"item_category_id,string"`
		IsVehicleWeapon     int               `json:"is_vehicle_weapon,string"`
		Name                map[string]string `json:"name"`
		Description         map[string]string `json:"description"`
		FactionID           Faction           `json:"faction_id,string"`
		MaxStackSize        int               `json:"max_stack_size,string"`
		ImageSetID          int64             `json:"image_set_id,string"`
		ImageID             int64             `json:"image_id,string"`
		ImagePath           string            `json:"image_path"`
		IsDefaultAttachment int               `json:"is_default_attachment,string"`
	}

	Profile *struct {
		ProfileTypeID          int64             `json:"profile_type_id,string"`
		ProfileTypeDescription string            `json:"profile_type_description"`
		FactionID              Faction           `json:"faction_id,string"`
		Name                   map[string]string `json:"name"`
		Description            map[string]string `json:"description"`
		ImageSetID             int64             `json:"image_set_id,string"`
		ImageID                int64             `json:"image_id,string"`
		ImagePath              string            `json:"image_path"`
		MovementSpeed          int               `json:"movement_speed,string"`
		BackpedalSpeedModifier float64           `json:"backpedal_speed_modifier,string"`
		SprintSpeedModifier    float64           `json:"sprint_speed_modifier,string"`
		StrafeSpeedModifier    float64           `json:"strafe_speed_modifier,string"`
	} `json:"profile"`

	Faction *struct {
		Name           map[string]string `json:"faction"`
		ImageSetID     int64             `json:"image_set_id,string"`
		ImageID        int64             `json:"image_id,string"`
		ImagePath      string            `json:"image_path"`
		CodeTag        string            `json:"code_tag"`
		UserSelectable int               `json:"user_selectable,string"`
	}

	OnlineStatus int `json:"online_status,string"`

	FriendList []struct {
		CharacterID   int64 `json:"character_id,string"`
		LastLoginTime int64 `json:"last_login_time,string"`
		Online        int   `json:"online,string"`
	} `json:"friend_list"`
}

func (g *Get) Character(search map[string]string, config *Config) (list []Character, err error) {
	err = g.Custom(&list, "character", search, config)
	if err != nil {
		return nil, err
	}

	return
}
