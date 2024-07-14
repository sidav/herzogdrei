package game

func (b *Battlefield) dealDamage(amount int, target Actor) {
	switch target.(type) {
	case *Commander:
		target.(*Commander).AsUnit.Health -= amount
	case *Unit:
		target.(*Unit).Health -= amount
	case *Building:
		target.(*Building).Hitpoints -= amount
	}
}
