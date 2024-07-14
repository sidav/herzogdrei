package game_static

type EffectStatic struct {
	SpriteCode      string
	DefaultLifeTime int
}

type EffectCode int

const (
	EFFECT_NONE = iota
	EFFECT_SMALL_EXPLOSION
	EFFECT_REGULAR_EXPLOSION
	EFFECT_BIGGER_EXPLOSION
)

var STableEffects = map[EffectCode]*EffectStatic{
	EFFECT_SMALL_EXPLOSION: {
		SpriteCode:      "smallexplosion",
		DefaultLifeTime: 16,
	},
	EFFECT_REGULAR_EXPLOSION: {
		SpriteCode:      "regularexplosion",
		DefaultLifeTime: 15,
	},
	EFFECT_BIGGER_EXPLOSION: {
		SpriteCode:      "biggerexplosion",
		DefaultLifeTime: 32,
	},
}
