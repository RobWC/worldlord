package mainmenu

import (
	"github.com/gtalent/starfish/gfx"
)

type MenuItem struct {
	text *gfx.Text
	x    int
	y    int
}

type MainMenu struct {
	Title      *gfx.Text
	logo       *gfx.Image
	anim       *gfx.Animation
	logoX      int
	logoY      int
	Cursor_mod bool
	menu       []MenuItem
}

func (mm *MainMenu) MoveLeft() {
	var speed_mod int
	speed_mod = 1
	if mm.Cursor_mod {
		speed_mod = 5
	}
	mm.logoY = mm.logoY - 16*speed_mod
}

func (mm *MainMenu) MoveRight() {
	var speed_mod int
	speed_mod = 1
	if mm.Cursor_mod {
		speed_mod = 5
	}
	mm.logoY = mm.logoY + 16*speed_mod
}

func (mm *MainMenu) MoveUp() {
	var speed_mod int
	speed_mod = 1
	if mm.Cursor_mod {
		speed_mod = 5
	}
	mm.logoX = mm.logoX - 16*speed_mod
}

func (mm *MainMenu) MoveDown() {
	var speed_mod int
	speed_mod = 1
	if mm.Cursor_mod {
		speed_mod = 5
	}
	mm.logoX = mm.logoX + 16*speed_mod
}

func (mm *MainMenu) Init() bool {
	return mm.init()
}

func (mm *MainMenu) init() bool {
	baseHeight := gfx.DisplayHeight()
	baseWidth := gfx.DisplayWidth()
	mm.logoY = baseHeight / 2
	mm.logoX = baseWidth / 2
	font := gfx.LoadFont("LiberationSans-Bold.ttf", 32)
	mm.logo = gfx.LoadImageSize("test.jpg", 16, 16)
	font.SetRGB(255, 60, 0)
	mm.menu = append(mm.menu, MenuItem{text: font.Write("Start"), x: (baseHeight / 2), y: (baseWidth / 5) + 50})
	mm.menu = append(mm.menu, MenuItem{text: font.Write("Do"), x: (baseHeight / 2), y: (baseWidth / 5) + 100})
	if font != nil {
		font.SetRGB(255, 60, 0)
		mm.Title = font.Write("Worldlord: Wrath of Chaos")
		font.Free()
	} else {
		println("Unable to load font")
		return false
	}
	return true
}

func (mm *MainMenu) Draw(c *gfx.Canvas) {
	baseHeight := gfx.DisplayHeight()
	baseWidth := gfx.DisplayWidth()
	c.SetRGB(0, 0, 0)
	c.FillRect(0, 0, baseWidth, baseHeight)
	//green square
	c.SetRGBA(0, 255, 100, 127)
	c.FillRect(0, 0, 100, 100)
	if mm.logoX > baseHeight-16 {
		mm.logoX = baseHeight - 16
	}
	if mm.logoY > baseWidth-16 {
		mm.logoY = baseWidth - 16
	}
	if mm.logoX < 0 {
		mm.logoX = 0
	}
	if mm.logoY < 0 {
		mm.logoY = 0
	}

	c.DrawText(mm.Title, baseHeight/4, baseWidth/5)
	for item := range mm.menu {
		c.DrawText(mm.menu[item].text, mm.menu[item].x, mm.menu[item].y)
	}
	c.PushViewport(42, 42, 500, 500)
	{
		//draw a green rect in a viewport
		c.SetRGBA(0, 255, 100, 127)
		c.FillRect(0, 0, 100, 100)
	}
	c.PopViewport()
	c.DrawImage(mm.logo, mm.logoY, mm.logoX)

}
