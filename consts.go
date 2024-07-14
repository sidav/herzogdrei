package main

const (
	WINDOW_W = 1366
	WINDOW_H = 768

	DESIRED_TPS = 60

	SPRITE_SCALE_FACTOR          = 4
	ORIGINAL_TILE_SIZE_IN_PIXELS = 16
	TILE_SIZE_IN_PIXELS          = ORIGINAL_TILE_SIZE_IN_PIXELS * SPRITE_SCALE_FACTOR
	PIXEL_TO_PHYSICAL_RATIO      = TILE_SIZE_IN_PIXELS

	zeroTiltColor     = 32
	strongerTiltColor = 128
)
