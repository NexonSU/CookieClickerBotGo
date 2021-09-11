package main

import (
	"github.com/NexonSU/CookieClickerBotGo/hotkeys"

	"golang.design/x/mainthread"
)

func main() {
	//robotgo.ScrollMouse(10, "down")
	//robotgo.MouseClick("left", true)
	//robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	//hotkeys
	mainthread.Init(hotkeys.Bind())
}
