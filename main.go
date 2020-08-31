package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	ticapp "github.com/mavincci/NilTOE/app"
	"github.com/mavincci/NilTOE/bit"
	"github.com/mavincci/NilTOE/ui"
)

func main() {
	var gameApplication = ticapp.Application{
		Game: bit.NewGame(),
	}
	application := app.New()
	window := application.NewWindow("NilTOE")

	window.Resize(fyne.NewSize( 320, 480))
	window.SetFixedSize(true)

	window.SetContent(ui.BuildUI(gameApplication.Game))

	window.ShowAndRun()
}
