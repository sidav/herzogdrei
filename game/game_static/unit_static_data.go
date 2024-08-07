package game_static

type UnitStatic struct {
	DisplayedName     string `json:"displayed_name"`
	ChassisSpriteCode string `json:"chassis_sprite_code,omitempty"`

	MaxHitpoints         int `json:"max_hitpoints,omitempty"`
	MaxFuel              int
	MovementSpeed        float64 `json:"movement_speed,omitempty"`
	ChassisRotationSpeed int     `json:"chassis_rotation_speed,omitempty"`

	IsCommander                        bool
	TransformsTo                       int
	TransformationAnimationSpritesCode string

	IsAircraft  bool `json:"is_aircraft,omitempty"`
	IsTransport bool `json:"is_transport,omitempty"`

	Cost      int `json:"cost,omitempty"`
	BuildTime int `json:"build_time,omitempty"` // seconds

	TurretsData []*TurretStatic

	OrderCosts map[int]int
}

func (us *UnitStatic) CanDoOrder(code int) bool {
	_, can := us.OrderCosts[code]
	return can
}

const (
	UNIT_COMPLANE = iota
	UNIT_COMROBOT
	UNIT_INFANTRY
	UNIT_RECON
	UNIT_QUAD
	UNIT_AATANK
	UNIT_TANK
	UNIT_DEVASTATOR
	UNIT_TURRET
)

func GetUnitStaticDataByCode(code int) *UnitStatic {
	return STableUnits[code]
}

var STableUnits = map[int]*UnitStatic{
	UNIT_COMPLANE: {
		DisplayedName:                      "C-VTOL",
		ChassisSpriteCode:                  "cplane",
		MaxHitpoints:                       100,
		MaxFuel:                            1000,
		MovementSpeed:                      0.16,
		ChassisRotationSpeed:               10,
		IsCommander:                        true,
		TransformsTo:                       UNIT_COMROBOT,
		TransformationAnimationSpritesCode: "cplanetransform",
		IsAircraft:                         true,
		IsTransport:                        true,
		TurretsData: []*TurretStatic{
			{
				AttacksLand:         false,
				FireRange:           6.5,
				FireSpreadDegrees:   5,
				ShotRangeSpread:     0.7,
				CooldownAfterVolley: 10,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "bullets",
					SplashRadius:              0.1,
					HitDamage:                 20,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.35,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_SMALL_EXPLOSION,
				},
			},
		},
	},
	UNIT_COMROBOT: {
		DisplayedName:                      "C-VTOL",
		ChassisSpriteCode:                  "crobot",
		MaxHitpoints:                       100,
		MaxFuel:                            1000,
		MovementSpeed:                      0.1,
		ChassisRotationSpeed:               5,
		Cost:                               10000,
		IsCommander:                        true,
		TransformsTo:                       UNIT_COMPLANE,
		TransformationAnimationSpritesCode: "crobottransform",
		IsAircraft:                         false,
		IsTransport:                        false,
		TurretsData: []*TurretStatic{
			{
				AttacksLand:         false,
				FireRange:           4,
				FireSpreadDegrees:   3,
				ShotRangeSpread:     0.4,
				CooldownAfterVolley: 25,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "shell",
					SplashRadius:              0.3,
					HitDamage:                 8,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.2,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_REGULAR_EXPLOSION,
				},
			},
		},
	},
	UNIT_INFANTRY: {
		DisplayedName:        "Infantry",
		ChassisSpriteCode:    "infantry",
		MaxHitpoints:         15,
		MovementSpeed:        0.016,
		ChassisRotationSpeed: 7,
		Cost:                 100,
		BuildTime:            7,
		OrderCosts: map[int]int{
			ORDER_CAPTURE: 150,
		},
	},
	UNIT_RECON: {
		DisplayedName:        "Recon",
		ChassisSpriteCode:    "infantry_recon",
		MaxHitpoints:         15,
		MovementSpeed:        0.027,
		ChassisRotationSpeed: 5,
		Cost:                 200,
		BuildTime:            10,
		OrderCosts: map[int]int{
			ORDER_CAPTURE: 150,
		},
	},
	UNIT_QUAD: {
		DisplayedName:        "Quad",
		ChassisSpriteCode:    "quad",
		MaxHitpoints:         25,
		MovementSpeed:        0.032,
		ChassisRotationSpeed: 4,
		Cost:                 200,
		BuildTime:            7,
		TurretsData: []*TurretStatic{
			{
				SpriteCode:          "",
				AttacksLand:         true,
				RotateSpeed:         0,
				FireRange:           5,
				FireSpreadDegrees:   15,
				ShotRangeSpread:     0.7,
				CooldownAfterVolley: 10,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "bullets",
					SplashRadius:              0,
					HitDamage:                 2,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.7,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_SMALL_EXPLOSION,
				},
			},
		},
		OrderCosts: map[int]int{
			ORDER_STANDBY:        0,
			ORDER_SEARCHNDESTROY: 150,
			ORDER_PATROL:         100,
		},
	},
	UNIT_AATANK: {
		DisplayedName:        "SAM",
		ChassisSpriteCode:    "tankchassis",
		MaxHitpoints:         25,
		MovementSpeed:        0.015,
		ChassisRotationSpeed: 2,
		Cost:                 100,
		BuildTime:            20,
		TurretsData: []*TurretStatic{
			{
				SpriteCode:          "aatankturret",
				AttacksLand:         false,
				AttacksAir:          true,
				RotateSpeed:         7,
				FireRange:           8,
				FireSpreadDegrees:   6,
				MaxShotsInVolley:    2,
				CooldownPerShot:     15,
				CooldownAfterVolley: 120,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "aamissile",
					SplashRadius:              0.25,
					HitDamage:                 20,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.125,
					RotationSpeed:             5,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_REGULAR_EXPLOSION,
				},
			},
		},
		OrderCosts: map[int]int{
			ORDER_STANDBY: 0,
			ORDER_PATROL:  100,
		},
	},
	UNIT_TANK: {
		DisplayedName:        "Tank",
		ChassisSpriteCode:    "tankchassis",
		MaxHitpoints:         55,
		MovementSpeed:        0.020,
		ChassisRotationSpeed: 3,
		Cost:                 100,
		BuildTime:            20,
		TurretsData: []*TurretStatic{
			{
				SpriteCode:          "tankcannon",
				AttacksLand:         true,
				RotateSpeed:         7,
				FireRange:           5,
				FireSpreadDegrees:   6,
				ShotRangeSpread:     0.7,
				CooldownAfterVolley: 55,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "shell",
					SplashRadius:              0.25,
					HitDamage:                 12,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.3,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_REGULAR_EXPLOSION,
				},
			},
		},
		OrderCosts: map[int]int{
			ORDER_STANDBY:        0,
			ORDER_SEARCHNDESTROY: 150,
			ORDER_PATROL:         100,
		},
	},
	UNIT_DEVASTATOR: {
		DisplayedName:        "Devastator",
		ChassisSpriteCode:    "devastator",
		MaxHitpoints:         250,
		MovementSpeed:        0.016,
		ChassisRotationSpeed: 2,
		Cost:                 200,
		BuildTime:            2,
		TurretsData: []*TurretStatic{
			{
				SpriteCode:          "devastatorturret",
				AttacksLand:         true,
				RotateSpeed:         0,
				FireRange:           5,
				FireSpreadDegrees:   6,
				ShotRangeSpread:     0.8,
				MaxShotsInVolley:    2,
				CooldownPerShot:     5,
				CooldownAfterVolley: 55,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "shell",
					SplashRadius:              0.25,
					HitDamage:                 11,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.3,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_REGULAR_EXPLOSION,
				},
			},
		},
		OrderCosts: map[int]int{
			ORDER_STANDBY:        0,
			ORDER_SEARCHNDESTROY: 150,
			ORDER_PATROL:         100,
		},
	},
	UNIT_TURRET: {
		DisplayedName:        "Turret",
		ChassisSpriteCode:    "turretunit",
		MaxHitpoints:         250,
		MovementSpeed:        0,
		ChassisRotationSpeed: 0,
		Cost:                 1000,
		BuildTime:            2,
		TurretsData: []*TurretStatic{
			{
				SpriteCode:          "minigunturret",
				AttacksLand:         true,
				RotateSpeed:         4,
				FireRange:           7.5,
				FireSpreadDegrees:   6,
				ShotRangeSpread:     0.8,
				MaxShotsInVolley:    1,
				CooldownPerShot:     0,
				CooldownAfterVolley: 6,
				FiredProjectileData: &ProjectileStatic{
					SpriteCode:                "bullets",
					SplashRadius:              0.15,
					HitDamage:                 4,
					SplashDamage:              10,
					Size:                      0.3,
					Speed:                     0.4,
					CreatesEffectOnImpact:     true,
					EffectCreatedOnImpactCode: EFFECT_SMALL_EXPLOSION,
				},
			},
		},
		OrderCosts: map[int]int{
			ORDER_STANDBY: 0,
		},
	},
}
