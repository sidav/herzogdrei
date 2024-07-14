package game

import "fmt"

type tile struct {
	Code               int
	landActorHere      Actor
	SpriteVariantIndex int
}

func (t *tile) GetStaticData() *tileStaticData {
	return STableTiles[t.Code]
}

func (t *tile) selectRandomSpriteVariant() {
	t.SpriteVariantIndex = rnd.Rand(len(t.GetStaticData().SpriteCodes))
}

const (
	TILE_SAND = iota
	TILE_BUILDABLE
	TILE_BUILDABLE_DAMAGED
	TILE_ROCK
	TILE_CONCRETE
)

type tileStaticData struct {
	SpriteCodes          []string
	BackgroundSpriteCode string
	CanBeWalkedOn        bool
}

func genNames(min, max int) []string {
	var n []string
	for i := min; i <= max; i++ {
		n = append(n, fmt.Sprintf("tile%d", i))
	}
	return n
}

var STableTiles = map[int]*tileStaticData{
	TILE_SAND: {
		SpriteCodes:   genNames(1, 16),
		CanBeWalkedOn: true,
	},
	TILE_ROCK: {
		SpriteCodes:          []string{"wall"},
		BackgroundSpriteCode: "tile1",
		CanBeWalkedOn:        false,
	},
}
