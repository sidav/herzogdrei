package game

import (
	. "herzog/game/game_static"
)

func (b *Battlefield) initFromStringMap() {
	b.Tiles = make([][]tile, len(gameStrMap[0]))
	for i := range b.Tiles {
		b.Tiles[i] = make([]tile, len(gameStrMap))
	}

	for i := range b.Tiles {
		for j := range b.Tiles[i] {
			b.Tiles[i][j].Code = TILE_SAND
		}
	}

	for y := range gameStrMap {
		for x := range gameStrMap[y] {
			switch rune(gameStrMap[y][x]) {
			case '#':
				b.Tiles[x][y].Code = TILE_ROCK
			case 'B':
				b.addActor(&Building{
					code:     BLD_BASE,
					TopLeftX: x,
					TopLeftY: y,
					Faction:  nil,
				})
			case '@':
				b.initAndPlaceNewFaction(x, y)
			}
		}
	}

	for i := range b.Tiles {
		for j := range b.Tiles[i] {
			b.Tiles[i][j].selectRandomSpriteVariant()
		}
	}
}

var gameStrMap = []string{
	"..............................###...............................",
	"..............................###...............................",
	"..@.................B.........###.......B..................@....",
	"..............................###...............................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"................................................................",
	"................................................................",
	"................................................................",
	"................................................................",
	".........................................B......................",
	"................................................................",
	"...........................................................#####",
	"#################..........B...................#################",
	"#####...........................................................",
	"................................................................",
	"..............B................................B................",
	"................................................................",
	"................................................................",
	"................................................................",
	".......................................B........................",
	"................................................................",
	"................................................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	"...............................#................................",
	".....................B.........#................................",
	".....@........................###............B..............@...",
	"..............................###...............................",
	"..............................###...............................",
	"..............................###...............................",
}
