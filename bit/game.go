package bit

import (
	"fmt"
)

type (
	Input	uint8
	State	struct {
		Player
		*Board
	}
	Game	struct {
		*State
		Error error
		Status string
		End bool
	}
)

func NewState() *State {
	return &State{ O, new(Board)}
}

func NewGame() *Game {
	game := new(Game)
	game.Reset()
	return game
}

func (game *Game) Reset() {
	game.End = false
	game.State = NewState()
	game.Board.Reset()
	game.Error = nil
	game.Status = ""
}

func (game *Game) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if game.Board.RowEqual(uint8(i)) && game.Board[i][0] != N {
			return true
		}
		if game.Board.ColEqual(uint8(i)) && game.Board[0][i] != N {
			return true
		}
	}
	if game.Board.DiagonalEqual(0) && game.Board[0][0] != N {
		return true
	}
	if game.Board.DiagonalEqual(1) && game.Board[2][0] != N {
		return true
	}

	return false
}

func (game *Game) CheckDraw() bool {
	for i := 0; i < 3; i++ {
		if game.RowContains(uint8(i), N) {
			return false
		}
	}
	return true
}

func (game *Game) nonErrorHalt(status string) {
	game.Error = nil
	game.Status = status
	game.End = true
}

func (game *Game) Next(input Input) {
	err := game.Move(input)
	if err != nil {
		game.Error = err
		return
	}

	game.Board[input / 3][input % 3] = game.Player

	if game.CheckWin() {
		game.nonErrorHalt(fmt.Sprintf("%s WON", game.State.Player))
		return
	} else if game.CheckDraw() {
		game.nonErrorHalt("DRAW")
		return
	}
	game.TogglePlayer()
}

func (game *Game) TogglePlayer() {
	if game.Error != nil {
		game.Error = nil
	}
	if game.Player == X {
		game.Player = O
	} else {
		game.Player = X
	}
}

func (game *Game) Move(input Input) error {
	if input > 8 || input < 0 {
		return INVALIDE_INPUT
	}
	if game.Board.Occupied(input) {
		return OCCUPIED_CELL
	}
	return nil
}
