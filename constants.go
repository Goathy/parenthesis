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

	emptyLine  string        = " "
	assocLeft  associativity = -1
	assocRight associativity = 1
)
