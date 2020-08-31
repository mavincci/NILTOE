package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/mavincci/NilTOE/bit"
)

type GamePanel struct {
	Turn	*widget.Label
	Status	*widget.Label
	Cell	[9]*widget.Button
}

func NewGamePanel() (panel *GamePanel) {
	panel = new(GamePanel)
	panel.Turn = widget.NewLabel("")
	panel.Status = widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{ false, true, false})
	panel.Status.Wrapping = fyne.TextWrapBreak

	for i := range panel.Cell {
		panel.Cell[i] = widget.NewButton("", nil)
	}
	return
}

func (panel *GamePanel) Handle(game *bit.Game, id uint8) func() {
	return func() {
		if !game.End {
			game.Next(bit.Input(id))
			panel.UpdateUI(game)
		}
		if game.End {
			for _, b := range panel.Cell {
				b.Disable()
			}
		}
	}
}

func (panel *GamePanel) UpdateUI(game *bit.Game) {
	panel.Turn.SetText("Turn: " + game.Player.String())
	if game.Error != nil {
		panel.Status.SetText(game.Error.Error())
	} else {
		panel.Status.SetText(game.Status)
	}
	for i := range panel.Cell {
		v := game.Board[i/3][i%3]
		panel.Cell[i].SetText(v.String())
	}
}
func (panel *GamePanel) BuildLabel(game *bit.Game) fyne.CanvasObject {
	l := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel(""),
		widget.NewLabelWithStyle("NilTOE",fyne.TextAlignCenter, fyne.TextStyle{ true, true, true }),
		widget.NewLabel(""),
		widget.NewLabel(""),
		panel.Status,
		panel.Turn,
		widget.NewButton("Reset", func() {
			for _, b := range panel.Cell {
				b.Enable()
			}
			game.Reset(); panel.UpdateUI(game) }),
	)
	return l
}

func (panel *GamePanel) BuildBoard(game *bit.Game) fyne.CanvasObject {
	b := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		)
	for i, btn := range panel.Cell {
		btn.OnTapped = panel.Handle(game, uint8(i))
		b.AddObject(btn)
	}
	return b
}

func BuildUI(game *bit.Game) fyne.CanvasObject {
	panel := NewGamePanel()
	panel.UpdateUI(game)
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(2),
		panel.BuildLabel(game),
		panel.BuildBoard(game),
		)
}
