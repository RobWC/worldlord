package main

import (
	"fmt"
	"github.com/gtalent/starfish/gfx"
	"github.com/gtalent/starfish/input"
	"lib/mainmenu"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	runtime.LockOSThread()

	if !gfx.OpenDisplay(1024, 768, false) {
		return
	}
	gfx.SetDisplayTitle("Worldlord: Wrath of Chaos")

	var pane mainmenu.MainMenu
	pane.Init()
	fmt.Println(pane)
	gfx.AddDrawer(&pane)

	quit := func() {
		gfx.CloseDisplay()
		os.Exit(0)
	}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			log.Println("Foos")
		}
	}()

	input.AddQuitFunc(quit)
	input.AddKeyPressFunc(func(e input.KeyEvent) {
		fmt.Println("Key Press ", e.Key)
		if e.Key == 1073742049 {
			pane.Cursor_mod = true
		}
		if e.Key == input.Key_Escape {
			quit()
		}
		if e.Key == input.Key_Down {
			pane.MoveDown()
		}
		if e.Key == input.Key_Up {
			pane.MoveUp()
		}
		if e.Key == input.Key_Left {
			pane.MoveLeft()
		}
		if e.Key == input.Key_Right {
			pane.MoveRight()
		}
	})
	input.AddKeyReleaseFunc(func(i input.KeyEvent) {
		fmt.Println("Key Release! ", i.Key)
		if i.Key == 1073742049 {
			pane.Cursor_mod = false
		}
	})

	log.Println("Start Worldlord")
	gfx.Main()
}
