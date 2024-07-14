package game_static

type BuildingStatic struct {
	W             int    `json:"w,omitempty"`
	H             int    `json:"h,omitempty"`
	DisplayedName string `json:"displayed_name,omitempty"`
	SpriteCode    string `json:"sprite_code,omitempty"`
	MaxHitpoints  int
	GivesMoney    int
}

const (
	BLD_MAIN_BASE = iota
	BLD_BASE
)

var STableBuildings = map[int]*BuildingStatic{
	BLD_MAIN_BASE: {
		W:             4,
		H:             4,
		DisplayedName: "HQ",
		SpriteCode:    "main2",
		GivesMoney:    4,
		MaxHitpoints:  1000,
	},
	BLD_BASE: {
		W:             2,
		H:             2,
		DisplayedName: "Base",
		SpriteCode:    "base",
		GivesMoney:    2,
	},
}
