package main

import (
	. "herzog/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	gameIsRunning bool
	battlefield   Battlefield
}

func (g *Game) Init() {
	g.gameIsRunning = false
	g.battlefield = Battlefield{}
	g.battlefield.Init()
}

func (g *Game) Start() {
	g.gameIsRunning = true
	r := renderer{}
	pc := playerController{}
	pc.init(g.battlefield.Factions[0])

	for rl.GetKeyPressed() != rl.KeyEscape {
		r.renderBattlefield(&g.battlefield)

		pc.control()

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
