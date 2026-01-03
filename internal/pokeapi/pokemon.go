package pokeapi

type Pokemon struct {
	Abilities              []Abilities     `json:"abilities"`
	BaseExperience         int             `json:"base_experience"`
	Cries                  Cries           `json:"cries"`
	Forms                  []Forms         `json:"forms"`
	GameIndices            []GameIndices   `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              []HeldItems     `json:"held_items"`
	ID                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []Moves         `json:"moves"`
	Name                   string          `json:"name"`
	Order                  int            `json:"order"`
	PastAbilities          []PastAbilities `json:"past_abilities"`
	PastTypes              []any           `json:"past_types"`
	Species                Species         `json:"species"`
	Stats                  []Stats         `json:"stats"`
	Types                  []Types         `json:"types"`
	Weight                 int             `json:"weight"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rarity  int     `json:"rarity"`
	Version Version `json:"version"`
}
type HeldItems struct {
	Item           Item             `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
}
type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PastAbilities struct {
	Abilities  []Abilities `json:"abilities"`
	Generation Generation  `json:"generation"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
