package game

func (b *Battlefield) DoEconomyTick() {
	for _, a := range b.Actors {
		if bld, ok := a.(*Building); ok {
			if bld.Faction != nil {
				bld.Faction.Gold += bld.GetStaticData().GivesMoney
			}
		}
	}

	for _, f := range b.Factions {
		if f.BuildsUnitNow {
			f.DoProductionStep()
		}
	}
}
