package bit

type (
	Player	uint8
	Board	[3][3]Player
)

const (
	X	Player = iota
	O
	N
)

var PSTR = map[Player]string {
	N: "", O: "O", X: "X",
}

func (board *Board) Reset() {
	for r := range board {
		for c := range board[r] {
			board[r][c] = N
		}
	}
}

func (board *Board) Occupied(input Input) bool {
	return board[input / 3][input % 3] != N
}

func (board *Board) RowEqual(row uint8) bool {
	return areEqual(board[row][0], board[row][1], board[row][2])
}

func (board *Board) ColEqual(col uint8) bool {
	return areEqual(board[0][col], board[1][col], board[2][col])
}

func (board *Board) DiagonalEqual(col uint8) bool {
	if col == 0 {
		return areEqual(board[0][0], board[1][1], board[2][2])
	}
	return areEqual(board[0][2], board[1][1], board[2][0])
}


func (p Player) String() string {
	return PSTR[p]
}

func (board *Board) RowContains(row uint8, player Player) bool {
	return rowContains(player, board[row][0], board[row][1], board[row][2])
}

func rowContains(player Player, things ...Player) bool {
	for _, v := range things {
		if v == player {
			return true
		}
	}
	return false
}

func areEqual(things ...interface{}) bool {
	t := things[0]
	for _, v := range things {
		if t != v {
			return false
		}
	}
	return true
}
