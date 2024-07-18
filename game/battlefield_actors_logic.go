package game

func (b *Battlefield) SwitchTilePointersForGroundActor(a Actor, oldtx, oldty, newtx, newty int) {
	if b.Tiles[oldtx][oldty].landActorHere != nil && b.Tiles[oldtx][oldty].landActorHere != a {
		panic("Error 1 on tile switch")
	}
	b.Tiles[oldtx][oldty].landActorHere = nil
	if b.Tiles[newtx][newty].landActorHere != nil && b.Tiles[newtx][newty].landActorHere != a {
		panic("Error 2 on tile switch")
	}
	b.Tiles[newtx][newty].landActorHere = a
}

// May be useful for AI or something
func (b *Battlefield) SelectActorWithHighestScore(getScore func(Actor) (score int, selectable bool)) Actor {
	var seletedActor Actor
	var selectedScore int
	for _, act := range b.Actors {
		score, actorSelectable := getScore(act)
		if !actorSelectable {
			continue
		}
		if seletedActor == nil || selectedScore < score {
			seletedActor = act
			selectedScore = score
		}
	}
	return seletedActor
}
