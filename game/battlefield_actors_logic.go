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
