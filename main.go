package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var clickerEnabled = true
var picClickerEnabled = true
var buyerEnabled = false

func main() {
	go Clicker()
	go PicClicker()
	go Buyer()
	BindHotkeys()
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

	robotgo.EventHook(hook.KeyDown, []string{"p", "ctrl", "shift"}, func(e hook.Event) {
		if picClickerEnabled {
			picClickerEnabled = false
			println("PicClicker disabled.")
		} else {
			picClickerEnabled = true
			println("PicClicker enabled.")
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
	for {
		time.Sleep(10 * time.Millisecond)
		if !clickerEnabled {
			continue
		}
		activeProcess, err := robotgo.FindName(robotgo.GetPID())
		if err != nil {
			continue
		}
		if activeProcess != "Cookie Clicker.exe" {
			continue
		}
		posX, posY, width, height := robotgo.GetBounds(robotgo.GetPID())
		if posX == 0 && posY == 0 && width == 0 && height == 0 {
			continue
		}
		for i := 0; i < 10; i++ {
			robotgo.MoveClick(int(float64(width)*0.155)+posX, int(float64(height)*0.42)+posY)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func PicClicker() {
	for {
		time.Sleep(100 * time.Millisecond)
		if !picClickerEnabled {
			continue
		}
		activeProcess, err := robotgo.FindName(robotgo.GetPID())
		if err != nil {
			continue
		}
		if activeProcess != "Cookie Clicker.exe" {
			continue
		}
		posX, posY, width, height := robotgo.GetBounds(robotgo.GetPID())
		if posX == 0 && posY == 0 && width == 0 && height == 0 {
			continue
		}
		windowScreen := robotgo.CaptureScreen(posX, posY, width, height)
		files, err := filepath.Glob("files/*.png")
		if err != nil {
			log.Println(err)
		}
		for _, f := range files {
			fx, fy := robotgo.FindPic(f, windowScreen, 0.1)
			if fx == -1 && fy == -1 {
				continue
			}
			robotgo.MoveClick(fx+posX, fy+posY)
		}
		robotgo.FreeBitmap(windowScreen)

	}
}

func Buyer() {
	for {
		time.Sleep(100 * time.Millisecond)
		if !buyerEnabled {
			continue
		}
		activeProcess, err := robotgo.FindName(robotgo.GetPID())
		if err != nil {
			continue
		}
		if activeProcess != "Cookie Clicker.exe" {
			continue
		}
		posX, posY, width, height := robotgo.GetBounds(robotgo.GetPID())
		if posX == 0 && posY == 0 && width == 0 && height == 0 {
			continue
		}
		fx, fy := robotgo.FindColorCS(robotgo.CHex(robotgo.RgbToHex(102, 255, 102)), posX+3*width/4, posY, width/4, height)
		if fx == -1 && fy == -1 {
			continue
		}
		robotgo.MoveMouse(fx+posX+3*width/4, fy+posY)
		robotgo.MouseClick("left")
	}
}
