package game

func (b *Battlefield) progressBuildingCapture(bld *Building, capturingFaction *Faction) {
	for i := range bld.CaptureProgress {
		if bld.CaptureProgress[i] != capturingFaction {
			bld.CaptureProgress[i] = capturingFaction
			break
		}
	}
	if bld.IsFullyCaptured() {
		bld.Faction = capturingFaction
	}
}
