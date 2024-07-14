package main

import (
	"fmt"
	satlas "herzog/lib/sprites_atlas"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	defaultFont rl.Font
	// index of array is faction color.
	tilesAtlaces       = map[string]*satlas.SpriteAtlas{}
	buildingsAtlaces   = map[string]*satlas.SpriteAtlas{}
	unitChassisAtlaces = map[string]*satlas.SpriteAtlas{}
	turretsAtlaces     = map[string]*satlas.SpriteAtlas{}
	projectilesAtlaces = map[string]*satlas.SpriteAtlas{}
	effectsAtlaces     = map[string]*satlas.SpriteAtlas{}

	uiAtlaces = map[string]*satlas.SpriteAtlas{}
)

func loadAssets() {
	defaultFont = rl.LoadFont("resources/flexi.ttf")
	satlas.SpriteScaleFactor = SPRITE_SCALE_FACTOR
	satlas.DebugWritef = debugWritef
	loadSprites()
}

func loadSprites() {
	tilesAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath := "resources/sprites/terrain/eric/"
	entries, err := os.ReadDir(currPath)
	if err != nil {
		panic("loading error")
	}
	for i, e := range entries {
		name := fmt.Sprintf("tile%d", i)
		tilesAtlaces[name] = satlas.CreateAtlasFromFile(currPath+e.Name(), 0, 0, 16, 16, 16, 16, 1, false, false)
	}
	currPath = "resources/sprites/terrain/"
	tilesAtlaces["wall"] = satlas.CreateAtlasFromFile(currPath+"wall.png", 0, 0, 16, 16, 16, 16, 1, false, false)
	tilesAtlaces["sand1"] = satlas.CreateAtlasFromFile(currPath+"sand1.png", 0, 0, 16, 16, 16, 16, 1, false, false)
	tilesAtlaces["sand2"] = satlas.CreateAtlasFromFile(currPath+"sand2.png", 0, 0, 16, 16, 16, 16, 1, false, false)
	tilesAtlaces["sand3"] = satlas.CreateAtlasFromFile(currPath+"sand3.png", 0, 0, 16, 16, 16, 16, 1, false, false)
	tilesAtlaces["rocks"] = satlas.CreateAtlasFromFile(currPath+"rocks.png", 0, 0, 16, 16, 16, 16, 1, false, false)

	buildingsAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/buildings/"
	buildingsAtlaces["main"] = satlas.CreateAtlasFromFile(currPath+"airfactory2.png", 0, 0, 96, 64, 48, 32, 1, false, true)
	buildingsAtlaces["main2"] = satlas.CreateAtlasFromFile(currPath+"HT_lab.png", 0, 0, 64, 64, 64, 64, 1, false, true)
	buildingsAtlaces["base"] = satlas.CreateAtlasFromFile(currPath+"base.png", 0, 0, 64, 64, 32, 32, 1, false, true)

	unitChassisAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/units/"
	unitChassisAtlaces["commander"] = satlas.CreateDirectionalAtlasFromFile(currPath+"aircrafts/command_plane.png", 32, 16, 1, 2, true)
	unitChassisAtlaces["infantry"] = satlas.CreateDirectionalAtlasFromFile(currPath+"infantry.png", 32, 16, 1, 2, true)
	unitChassisAtlaces["infantry_recon"] = satlas.CreateDirectionalAtlasFromFile(currPath+"infantry_recon.png", 32, 16, 1, 2, true)
	unitChassisAtlaces["quad"] = satlas.CreateDirectionalAtlasFromFile(currPath+"quad.png", 32, 16, 1, 2, true)
	unitChassisAtlaces["tankchassis"] = satlas.CreateDirectionalAtlasFromFile(currPath+"tank_chassis.png", 32, 16, 1, 2, true)
	unitChassisAtlaces["devastator"] = satlas.CreateDirectionalAtlasFromFile(currPath+"devastator.png", 26, 16, 1, 2, true)

	turretsAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/units/"
	turretsAtlaces["tankcannon"] = satlas.CreateDirectionalAtlasFromFile(currPath+"tank_cannon.png", 32, 16, 1, 2, true)
	turretsAtlaces["aatankturret"] = satlas.CreateDirectionalAtlasFromFile(currPath+"aatank_turret.png", 32, 16, 1, 2, true)
	turretsAtlaces["devastatorturret"] = satlas.CreateDirectionalAtlasFromFile(currPath+"devastator_turret.png", 26, 16, 1, 2, true)

	projectilesAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/projectiles/"
	projectilesAtlaces["shell"] = satlas.CreateDirectionalAtlasFromFile(currPath+"shell.png", 32, 16, 1, 2, false)
	projectilesAtlaces["bullets"] = satlas.CreateDirectionalAtlasFromFile(currPath+"bullets.png", 32, 8, 1, 2, false)
	projectilesAtlaces["missile"] = satlas.CreateDirectionalAtlasFromFile(currPath+"missile.png", 32, 16, 1, 2, false)
	projectilesAtlaces["aamissile"] = satlas.CreateDirectionalAtlasFromFile(currPath+"aamissile.png", 32, 16, 1, 2, false)
	projectilesAtlaces["omni"] = satlas.CreateDirectionalAtlasFromFile(currPath+"omni.png", 32, 16, 1, 2, false)

	uiAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/ui/"
	uiAtlaces["factionflag"] = satlas.CreateAtlasFromFile(currPath+"building_faction_flag.png", 0, 0, 4, 4, 4, 4, 4, false, true)

	effectsAtlaces = make(map[string]*satlas.SpriteAtlas)
	currPath = "resources/sprites/effects/"
	effectsAtlaces["smallexplosion"] = satlas.CreateAtlasFromFile(currPath+"explosion_small.png", 0, 0, 4, 4, 4, 4, 16, false, false)
	effectsAtlaces["regularexplosion"] = satlas.CreateAtlasFromFile(currPath+"explosion.png", 0, 0, 16, 16, 16, 16, 3, false, false)
	effectsAtlaces["biggerexplosion"] = satlas.CreateAtlasFromFile(currPath+"explosion_bigger.png", 0, 0, 40, 40, 20, 20, 3, false, false)
}
