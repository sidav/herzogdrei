package game

import "herzog/game/game_static"

func (b *Battlefield) CleanProjectiles() {
	for i := len(b.Projectiles) - 1; i >= 0; i-- {
		if b.Projectiles[i].SetToRemove {
			b.Projectiles = append(b.Projectiles[:i], b.Projectiles[i+1:]...)
		}
	}
}

func (b *Battlefield) CleanEffects() {
	for i := len(b.Effects) - 1; i >= 0; i-- {
		if b.Effects[i].GetExpirationPercent(b.CurrentTick) > 100 {
			b.Effects = append(b.Effects[:i], b.Effects[i+1:]...)
		}
	}
}

func (b *Battlefield) CleanDeadUnits() {
	for i := len(b.Actors) - 1; i >= 0; i-- {
		if unt, ok := b.Actors[i].(*Unit); ok {
			if !unt.IsAlive() {
				b.Actors = append(b.Actors[:i], b.Actors[i+1:]...)
				b.Effects = append(b.Effects, &Effect{
					CenterX:            unt.CenterX,
					CenterY:            unt.CenterY,
					Code:               game_static.EFFECT_BIGGER_EXPLOSION,
					CreationTick:       b.CurrentTick,
					SplashCircleRadius: 1.5,
				})

				// clean tile
				tx, ty := unt.GetTileCoords()
				cleared := false // debug, should be safe to remove
				for x := tx - 1; x <= tx+1; x++ {
					for y := ty - 1; y <= ty+1; y++ {
						if b.GetGroundActorAtTileCoordinates(x, y) == unt {
							cleared = true
							b.Tiles[x][y].landActorHere = nil
						}
					}
				}
				if !cleared {
					panic("unsuccessful tile clean!")
				}
			}
		}
	}
}
