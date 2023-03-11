package parenthesis

type associativity int

const (
	opPow      string = "^"
	opMulti    string = "*"
	opDiv      string = "/"
	opAdd      string = "+"
	opSub      string = "-"
	opLeftPar  string = "("
	opRightPar string = ")"

	blank string = " "
	empty string = ""

	assocLeft  associativity = -1
	assocRight associativity = 1
)
