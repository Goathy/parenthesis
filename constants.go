package parenthesis

type associativity int

const (
	OpPow      string = "^"
	OpMulti    string = "*"
	OpDiv      string = "/"
	OpAdd      string = "+"
	OpSub      string = "-"
	OpLeftPar  string = "("
	OpRightPar string = ")"

	Blank string = " "
	Empty string = ""

	assocLeft  associativity = -1
	assocRight associativity = 1
)
