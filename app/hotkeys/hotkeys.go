package hotkeys

import (
	"context"

	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
)

func Bind() {
	// Register a desired hotkey.
	hk, err := hotkey.Register([]hotkey.Modifier{hotkey.ModAlt}, hotkey.KeyZ)
	if err != nil {
		robotgo.ShowAlert("Error", err.Error(), "OK", "")
	}

	// Start listen hotkey event whenever you feel it is ready.
	triggered := hk.Listen(context.Background())
	for range triggered {
		println("hotkey ctrl+s is triggered")
	}
}
