package bit

type InputError uint8

const (
	OCCUPIED_CELL	InputError = iota
	INVALIDE_INPUT
)

var ESTR = map[InputError] string {
	OCCUPIED_CELL: "OCCUPIED_CELL",
	INVALIDE_INPUT: "INVALIDE_INPUT",
}

func (ie InputError) Error() string {
	return "InputError: " + ESTR[ie]
}
