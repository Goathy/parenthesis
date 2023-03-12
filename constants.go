package parenthesis

type Associativity int

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

	AssocLeft  Associativity = -1
	AssocRight Associativity = 1
)
