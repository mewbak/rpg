package rpg

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func InventoryLoop(win *pixelgl.Window, world *World) {
	x, y := win.Bounds().Center().XY()
	text := NewTextSmooth(14)
	text.WriteString("\tGAME PAUSED\n\n")
	text.WriteString("\tESC or any key to return\n\n")
	text.WriteString(world.Char.Stats.String())
	text.WriteString(fmt.Sprintf("Level %v\nHealth: %v\nMana: %v\nXP: %v/%v", world.Char.Level, world.Char.Health, world.Char.Mana, world.Char.Stats.XP, world.Char.NextLevel()))
	for _, item := range world.Char.Inventory {
		if item.Effect != nil {
			text.WriteString(fmt.Sprintf("\n Effects: %q", item.Name))

		}
	}

	text.WriteString("\n\n===INVENTORY===\n" + FormatItemList(world.Char.Inventory))
	for !win.Closed() {

		win.Clear(colornames.Black)
		x++
		y++
		if x > win.Bounds().Max.X {
			x = 0
		}
		if y > win.Bounds().Max.Y {
			y = 0
		}
		imd := imdraw.New(nil)
		imd.Color = colornames.Green
		imd.Push(pixel.V(0, 0), win.Bounds().Max)
		imd.Rectangle(4)
		imd.Color = colornames.White
		imd.Push(pixel.V(x+100, y+100), pixel.V(x+110, y+110))
		imd.Rectangle(0)
		imd.Push(pixel.V(x, y), pixel.V(x+10, y+10))
		imd.Rectangle(0)
		imd.Draw(win)

		if win.Typed() != "" || win.JustPressed(pixelgl.KeyEscape) {
			break
		}
		text.Draw(win, pixel.IM.Moved(pixel.V(30, 500)))
		win.SetTitle("AERPG inventory")
		win.Update()
	}
}
