package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// defer recoverPanicToFile()
	rl.InitWindow(int32(WINDOW_W), int32(WINDOW_H), "DAS IST KEIN HERZOG ZWEI!")
	rl.SetTargetFPS(DESIRED_TPS)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetExitKey(rl.KeyF10)

	loadAssets()

	game := Game{}
	game.Init()
	game.Start()

	rl.CloseWindow()
}
