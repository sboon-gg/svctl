package prbf2

type Murmur struct {
	Port int `yaml:"port"`
}

type Ports struct {
	Game    int `yaml:"game"`
	Stats   int `yaml:"stats"`
	Gamespy int `yaml:"gamespy"`
	Console int `yaml:"console"`
	Prism   int `yaml:"prism"`
}

type SponsoreMessage struct {
	Enabled  bool   `yaml:"enabled"`
	Message  string `yaml:"message"`
	Interval int    `yaml:"interval"`
}

type Squads struct {
	NoSquadsBefore       int     `yaml:"noSquadsBefore"`
	ResignEarly          bool    `yaml:"resignEarly"`
	KickLimit            int     `yaml:"kickLimit"`
	KickSquadLess        bool    `yaml:"kickSquadLess"`
	KickSquadLessTime    int     `yaml:"kickSquadLessTime"`
	KickSquadLessAFK     bool    `yaml:"kickSquadLessAFK"`
	KickSquadLessAFKTime int     `yaml:"kickSquadLessAFKTime"`
	KickAFKPercent       float64 `yaml:"kickAFKPercent"`
	KickSquadedAFK       bool    `yaml:"kickSquadedAFK"`
	KickSquadedAFKTime   int     `yaml:"kickSquadedAFKTime"`
}

type Admin struct {
	Name  string `yaml:"name"`
	Hash  string `yaml:"hash"`
	Level int    `yaml:"level"`
}

type PrismUser struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Level    int    `yaml:"level"`
}

type Maplist struct {
	Name  string `yaml:"name,omitempty"`
	Layer string `yaml:"layer,omitempty"`
	Mode  string `yaml:"mode"`
}

type CommandLevels struct {
	Aa           int `yaml:"aa"`
	Ab           int `yaml:"ab"`
	Admins       int `yaml:"admins"`
	Assignlead   int `yaml:"assignlead"`
	Br           int `yaml:"br"`
	Ban          int `yaml:"ban"`
	Banid        int `yaml:"banid"`
	Ec           int `yaml:"ec"`
	Flip         int `yaml:"flip"`
	Fly          int `yaml:"fly"`
	Kick         int `yaml:"kick"`
	Kill         int `yaml:"kill"`
	Givelead     int `yaml:"givelead"`
	Hash         int `yaml:"hash"`
	Help         int `yaml:"help"`
	History      int `yaml:"history"`
	Init         int `yaml:"init"`
	Info         int `yaml:"info"`
	Mapvote      int `yaml:"mapvote"`
	Message      int `yaml:"message"`
	Roundban     int `yaml:"roundban"`
	Reload       int `yaml:"reload"`
	Report       int `yaml:"report"`
	Resetsquads  int `yaml:"resetsquads"`
	Resign       int `yaml:"resign"`
	Resignall    int `yaml:"resignall"`
	Reportplayer int `yaml:"reportplayer"`
	Runnext      int `yaml:"runnext"`
	Rules        int `yaml:"rules"`
	Say          int `yaml:"say"`
	Sayteam      int `yaml:"sayteam"`
	Scramble     int `yaml:"scramble"`
	Setnext      int `yaml:"setnext"`
	Showafk      int `yaml:"showafk"`
	Shownext     int `yaml:"shownext"`
	Stopserver   int `yaml:"stopserver"`
	Swapteams    int `yaml:"swapteams"`
	Switch       int `yaml:"switch"`
	Timeban      int `yaml:"timeban"`
	Timebanid    int `yaml:"timebanid"`
	Tempban      int `yaml:"tempban"`
	Tickets      int `yaml:"tickets"`
	Unban        int `yaml:"unban"`
	Unbanid      int `yaml:"unbanid"`
	Unbanname    int `yaml:"unbanname"`
	Ungrief      int `yaml:"ungrief"`
	Warn         int `yaml:"warn"`
	Website      int `yaml:"website"`
}

type PRBF2 struct {
	Ports            Ports               `yaml:"ports"`
	Countryflag      string              `yaml:"countryflag"`
	IP               string              `yaml:"ip"`
	ExternalIP       string              `yaml:"externalIP"`
	Name             string              `yaml:"name"`
	Password         string              `yaml:"password"`
	Internet         int                 `yaml:"internet"`
	SponsorLogoURL   string              `yaml:"sponsorLogoURL"`
	CommunityLogoURL string              `yaml:"communityLogoURL"`
	MaxPlayers       int                 `yaml:"maxPlayers"`
	VotingEnabled    int                 `yaml:"votingEnabled"`
	DemoQuality      int                 `yaml:"demoQuality"`
	SponsoreMessage  SponsoreMessage     `yaml:"sponsoreMessage"`
	Squads           Squads              `yaml:"squads"`
	Admins           []Admin             `yaml:"admins"`
	LiteAdmins       []Admin             `yaml:"liteAdmins"`
	PrismUsers       []PrismUser         `yaml:"prismUsers"`
	ReservedSlotsNum int                 `yaml:"reservedSlotsNum"`
	ReservedSlots    []string            `yaml:"reservedSlots"`
	Maplist          []Maplist           `yaml:"maplist"`
	AllMaps          []interface{}       `yaml:"allMaps"`
	CommandAliases   map[string][]string `yaml:"commandAliases"`
	Reasons          map[string]string   `yaml:"reasons"`
	CommandLevels    CommandLevels       `yaml:"commandLevels"`
}
