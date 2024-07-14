package game_static

type ProjectileStatic struct {
	SpriteCode                string     `json:"sprite_code,omitempty"`
	Size                      float64    `json:"size,omitempty"`
	Speed                     float64    `json:"speed,omitempty"`
	SplashRadius              float64    `json:"splash_radius,omitempty"`
	CreatesEffectOnImpact     bool       `json:"creates_effect_on_impact,omitempty"`
	EffectCreatedOnImpactCode EffectCode `json:"effect_created_on_impact_code,omitempty"`
	RotationSpeed             int        `json:"rotation_speed,omitempty"`

	HitDamage    int `json:"hit_damage,omitempty"`
	SplashDamage int `json:"splash_damage,omitempty"`
}
