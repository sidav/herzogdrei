package game

import (
	. "herzog/game/game_static"
)

func (b *Battlefield) CreateCommanderForFaction(faction *Faction, x, y float64) {
	com := &Commander{
		AsUnit: *b.CreateNewUnit(UNIT_COMPLANE, faction, x, y),
	}
	b.addActor(com)
	faction.Commander = com
}

func (b *Battlefield) CreateNewUnit(code int, faction *Faction, x, y float64) *Unit {
	unt := &Unit{
		Code:    code,
		Faction: faction,
		CenterX: x,
		CenterY: y,
	}
	unt.Health = unt.GetStaticData().MaxHitpoints
	unt.Fuel = unt.GetStaticData().MaxFuel

	sd := unt.GetStaticData()
	if sd.TurretsData != nil {
		for i := range sd.TurretsData {
			unt.Turrets = append(unt.Turrets,
				&Turret{staticData: sd.TurretsData[i], RotationDegree: 270},
			)
		}
		unt.snapTurretsDegreesToChassis()
	}

	return unt
}

func (b *Battlefield) AddNewEffect(effCode EffectCode, x, y, radius float64) {
	b.Effects = append(b.Effects,
		&Effect{
			CenterX:            x,
			CenterY:            y,
			Code:               effCode,
			SplashCircleRadius: radius,
			CreationTick:       b.CurrentTick,
		})
}
