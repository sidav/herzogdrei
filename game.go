package main

import (
	"herzog/ai"
	. "herzog/game"
	"herzog/lib/random/pcgrandom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	gameIsRunning bool
	battlefield   Battlefield
}

func (g *Game) Init() {
	g.gameIsRunning = false
	g.battlefield = Battlefield{}
}

func (g *Game) Start() {
	rnd := pcgrandom.New(-1)
	g.battlefield.Init(rnd)
	g.gameIsRunning = true
	r := renderer{}
	pc := playerController{}

	pc.init(g.battlefield.Factions[0])
	var ais []*ai.AiStruct
	for i := 1; i < len(g.battlefield.Factions); i++ {
		ais = append(ais, ai.CreateNewAiStruct(&g.battlefield, g.battlefield.Factions[i]))
	}
	ai.SetPRNG(rnd)

	for rl.GetKeyPressed() != rl.KeyEscape {
		r.renderBattlefield(&g.battlefield)

		pc.control()
		for i := range ais {
			ais[i].Act()
		}

		for _, a := range g.battlefield.Actors {
			switch a.(type) {
			case *Commander:
				com := a.(*Commander)
				g.battlefield.ExecuteCommanderAction(com)

			case *Unit:
				unt := a.(*Unit)
				// Unit should finish the action before thinking of the next one
				if unt.Action.Kind == ACTION_NONE {
					g.battlefield.ExecuteUnitOrder(unt)
				}
				g.battlefield.ExecuteUnitAction(unt)
				g.battlefield.ActForAllActorsTurrets(unt)
			}
		}

		for _, p := range g.battlefield.Projectiles {
			g.battlefield.ActForProjectile(p)
		}

		if g.battlefield.CurrentTick%(DESIRED_TPS/2) == 0 {
			g.battlefield.DoEconomyTick()
		}

		if g.battlefield.CurrentTick%(DESIRED_TPS/4) == 0 {
			g.battlefield.CleanProjectiles()
		}
		if g.battlefield.CurrentTick%(DESIRED_TPS/4) == 1 {
			g.battlefield.CleanDeadUnits()
		}
		if g.battlefield.CurrentTick%(DESIRED_TPS/4) == 2 {
			g.battlefield.CleanEffects()
		}

		g.battlefield.CurrentTick++
	}
}
