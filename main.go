package main

import (
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var clickerEnabled = true
var goldClickerEnabled = true
var buyerEnabled = true

func main() {
	go BindHotkeys()
	go Clicker()
	go GoldClicker()
	go Buyer()

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

func BindHotkeys() {
	robotgo.EventHook(hook.KeyDown, []string{"c", "ctrl", "shift"}, func(e hook.Event) {
		if clickerEnabled {
			clickerEnabled = false
			println("Clicker disabled.")
		} else {
			clickerEnabled = true
			println("Clicker enabled.")
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"g", "ctrl", "shift"}, func(e hook.Event) {
		if goldClickerEnabled {
			goldClickerEnabled = false
			println("GoldClicker disabled.")
		} else {
			goldClickerEnabled = true
			println("GoldClicker enabled.")
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"b", "ctrl", "shift"}, func(e hook.Event) {
		if buyerEnabled {
			buyerEnabled = false
			println("Buyer disabled.")
		} else {
			buyerEnabled = true
			println("Buyer enabled.")
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		os.Exit(0)
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func Clicker() {
	posX := 0
	posY := 0
	width := 0
	height := 0
	for {
		posX, posY, width, height = robotgo.GetBounds(robotgo.GetPID())
		if strings.Contains(robotgo.GetTitle(), "Cookie Clicker") && clickerEnabled && posX != 0 && posY != 0 && width != 0 && height != 0 {
			robotgo.MoveMouse(int(float64(width)*0.155)+posX, int(float64(height)*0.42)+posY)
			robotgo.MouseClick("left")
			time.Sleep(2 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func GoldClicker() {
	posX := 0
	posY := 0
	width := 0
	height := 0
	fx := 0
	fy := 0
	windowScreen := robotgo.CaptureScreen()
	for {
		posX, posY, width, height = robotgo.GetBounds(robotgo.GetPID())
		if strings.Contains(robotgo.GetTitle(), "Cookie Clicker") && goldClickerEnabled && posX != 0 && posY != 0 && width != 0 && height != 0 {
			windowScreen = robotgo.CaptureScreen(posX, posY, width, height)
			fx, fy = robotgo.FindPic("goldCookie.png", windowScreen, 0.1)
			if fx != -1 && fy != -1 {
				robotgo.MoveMouse(fx+posX, fy+posY)
				robotgo.MouseClick("left")
			}
			robotgo.FreeBitmap(windowScreen)
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func Buyer() {
	posX := 0
	posY := 0
	width := 0
	height := 0
	fx := 0
	fy := 0
	for {
		posX, posY, width, height = robotgo.GetBounds(robotgo.GetPID())
		if strings.Contains(robotgo.GetTitle(), "Cookie Clicker") && buyerEnabled && posX != 0 && posY != 0 && width != 0 && height != 0 {
			fx, fy = robotgo.FindColorCS(robotgo.CHex(robotgo.RgbToHex(102, 255, 102)), posX+3*width/4, posY+height/2, width/4, height/2, 0.1)
			if fx != -1 && fy != -1 {
				robotgo.MoveMouse(fx+posX+3*width/4, fy+posY+height/2)
				robotgo.MouseClick("left")
			}
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
