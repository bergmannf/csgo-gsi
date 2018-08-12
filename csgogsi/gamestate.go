package csgogsi

import ()

type GSIData struct {
	Bomb       Bomb              `json:"bomb"`
	Provider   Provider          `json:"provider"`
	Map        Map               `json:"map"`
	Round      Round             `json:"round"`
	Player     Player            `json:"player"`
	AllPlayers map[string]Player `json:"allplayers"`
}

type Bomb struct {
	State    string `json:"state"`
	Position string `json:"position"`
	Player   int    `json:"player"`
}

type Provider struct {
	Name      string `json:"name"`
	Appid     int    `json:"appid"`
	Version   int    `json:"version"`
	Steamid   string `json:"steamid"`
	Timestamp int64  `json:"timestamp"`
}

type Map struct {
	Mode                  string            `json:"mode"`
	Name                  string            `json:"name"`
	Phase                 string            `json:"phase"`
	Round                 int               `json:"round"`
	TeamCT                Team              `json:"team_ct"`
	TeamT                 Team              `json:"team_t"`
	NumMatchesToWinSeries int               `json:"num_matches_to_win_series"`
	CurrentSpectators     int               `json:"current_spectators"`
	SouvenirsTotal        int               `json:"souvenirs_total"`
	RoundWins             map[string]string `json:"round_wins"`
}

type Round struct {
	Phase string `json:"phase"`
}

type Team struct {
	Score                int `json:"score"`
	TimeoutsRemaining    int `json:"timeouts_remaining"`
	MatchesWonThisSeries int `json:"matches_won_this_series"`
}

type Player struct {
	Steamid      string            `json:"steamid"`
	Clan         string            `json:"clan"`
	Name         string            `json:"name"`
	ObserverSlot int               `json:"observer_slot"`
	Team         string            `json:"team"`
	Activity     string            `json:"activity"`
	State        PlayerState       `json:"state"`
	MatchStats   MatchStats        `json:"match_stats"`
	Weapons      map[string]Weapon `json:"weapons"`
	Position     string            `json:"position"`
	Forward      string            `json:"forward"`
}

type PlayerState struct {
	Health      int  `json:"health"`
	Armor       int  `json:"armor"`
	Helmet      bool `json:"helmet"`
	Defusekit   bool `json:"defusekit"`
	Flashed     int  `json:"flashed"`
	Smoked      int  `json:"smoked"`
	Burning     int  `json:"burning"`
	Money       int  `json:"money"`
	RoundKills  int  `json:"round_kills"`
	RoundKillhs int  `json:"round_killshs"`
	EquipValue  int  `json:"equip_value"`
}

type Weapon struct {
	Name        string `json:"name"`
	Paintkit    string `json:"paintkit"`
	Type        string `json:"type"`
	State       string `json:"state"`
	AmmoClip    int    `json:"ammo_clip"`
	AmmoClipMax int    `json:"ammo_clip_max"`
	AmmoReserve int    `json:"ammo_reserve"`
}

type MatchStats struct {
	Kills   int `json:"kills"`
	Assists int `json:"assists"`
	Deaths  int `json:"deaths"`
	Mvps    int `json:"mvps"`
	Score   int `json:"score"`
}
