package csgogsi

import ()

type GSIData struct {
	Bomb       Bomb
	Provider   Provider
	Map        Map
	Round      Round
	Player     Player
	AllPlayers map[string]Player
}

type Bomb struct {
	State    string
	Position string
	Player   int
}

type Provider struct {
	Name      string
	Appid     int
	Version   int
	Steamid   int
	Timestamp int
}

type Map struct {
	Mode                  string
	Name                  string
	Phase                 string
	Round                 int
	TeamCT                Team              `json:"team_ct"`
	TeamT                 Team              `json:"team_t"`
	NumMatchesToWinSeries int               `json:"num_matches_to_win_series"`
	CurrentSpectators     int               `json:"current_spectators"`
	SouvenirsTotal        int               `json:"souvenirs_total"`
	RoundWins             map[string]string `json:"round_wins"`
}

type Round struct {
	Phase string
}

type Team struct {
	Score                int
	TimeoutsRemaining    int `json:"timeouts_remaining"`
	MatchesWonThisSeries int `json:"matches_won_this_series"`
}

type Player struct {
	Steamid      string
	Clan         string
	Name         string
	ObserverSlot int `json:"observer_slot"`
	Team         string
	Activity     string
	State        PlayerState
	MatchStats   MatchStats `json:"match_stats"`
	Weapons      map[string]Weapon
	Position     string
	Forward      string
}

type PlayerState struct {
	Health      int
	Armor       int
	Helmet      bool
	Defusekit   bool
	Flashed     int
	Smoked      int
	Burning     int
	Money       int
	RoundKills  int `json:"round_kills"`
	RoundKillhs int `json:"round_kills_hs"`
	EquipValue  int
}

type Weapon struct {
	Name        string
	Paintkit    string
	Type        string //!
	State       string
	AmmoClip    int `json:"ammo_clip"`
	AmmoClipMax int `json:"ammo_clip_max"`
	AmmoReserve int `json:"ammo_reserve"`
}

type MatchStats struct {
	Kills   int
	Assists int
	Deaths  int
	Mvps    int
	Score   int
}

type GSIConfiguration struct {
	Name                 string
	Uri                  string
	Timeout              string
	Buffer               string
	Throttle             string
	Heartbeat            string
	Provider             string
	Map                  string
	Round                string
	PlayerId             string `json:"player_id"`
	AllplayersId         string `json:"allplayers_id"`
	PlayerState          string `json:"player_state"`
	AllplayerState       string `json:"allplayer_state"`
	AllplayersMatchStats string `json:"allplayers_match_stats"`
	AllplayersWeapon     string `json:"allplayers_weapon"`
	AllplayersPosition   string `json:"allplayers_position"`
	PhaseCountdowns      string `json:"phase_countdowns"`
}
