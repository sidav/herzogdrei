package spritesatlas

import "image/color"

var zeroTiltColor uint8 = 32
var strongerTiltColor uint8 = 128

var FactionColors = []color.RGBA{
	{
		R: zeroTiltColor,
		G: zeroTiltColor,
		B: 255,
		A: 255,
	},
	{
		R: 255,
		G: zeroTiltColor,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: zeroTiltColor,
		G: 255,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: 255,
		G: 255,
		B: zeroTiltColor,
		A: 255,
	},
	{
		R: zeroTiltColor,
		G: 255,
		B: 255,
		A: 255,
	},
	{
		R: 255,
		G: zeroTiltColor,
		B: 255,
		A: 255,
	},
}
