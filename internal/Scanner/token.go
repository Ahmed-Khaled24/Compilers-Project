package Scanner

type BASETOKENTYPE int64
type ComparsionToken int64
type ArithmeticToken int64
type BracketToken int64

const (
	ERROR BASETOKENTYPE = iota
	RESERVED
	SPECIALSYMBOL
	ASSIGNMENT
	NUMBER
	IDENTIFIER
)
const (
	PLUS ArithmeticToken = iota
	MINUS
	MULT
	DIV
)
const (
	EQ ComparsionToken = iota
	LT
	GT
	LTE
	GTE
)
const (
	OPENBRACKET BracketToken = iota
	CLOSEDBRACKET
	OPENBRACE
	CLOSEDBRACE
)

func (t ComparsionToken) String() string {
	switch t {
	case EQ:
		return "EQUAL"
	case LT:
		return "LESSTHAN"
	case GT:
		return "GREATERTHAN"
	case LTE:
		return "LESSTHANOREQUAL"
	case GTE:
		return "GREATERTHANOREQUAL"
	}

	return "UNKNOW"
}
func (t BracketToken) String() string {
	switch t {
	case OPENBRACKET:
		return "OPENBRACKET"
	case CLOSEDBRACKET:
		return "CLOSEDBRACKET"
	case OPENBRACE:
		return "OPENBRACE"
	case CLOSEDBRACE:
		return "CLOSEDBRACE"
	}
	return "UNKNOW"
}
func (t ArithmeticToken) String() string {
	switch t {
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULT:
		return "MULT"
	case DIV:
		return "DIV"
	}
	return "UNKNOW"
}

func (t BASETOKENTYPE) String() string {
	switch t {
	case ERROR:
		return "ERROR"
	case RESERVED:
		return "RESERVED"
	case SPECIALSYMBOL:
		return "SPECIALSYMBOL"
	case ASSIGNMENT:
		return "ASSIGNMENT"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	}
	return "UNKNOW"
}

func get_child_attribute(base_type BASETOKENTYPE, token_val string) string {
	switch base_type {
	case SPECIALSYMBOL:
		return handle_special_symbols(token_val)

	}
	return "UNKOWN"
}

func handle_special_symbols(token_val string) string {
	switch token_val {
	case "<":
		return LT.String()
	case "=":
		return EQ.String()
	case ">":
		return GT.String()
	case "<=":
		return LTE.String()
	case ">=":
		return GT.String()
	}
	return "UNKNOWN"
}

type Token struct {
	token_base_type BASETOKENTYPE
	token_value     string
	token_type      string
}

func CreateToken(base_type BASETOKENTYPE, token_val string) *Token {
	// naming convision: The method or function you want to access should start with an uppercase letter in Go
	// constructor for Token struct
	// param:
	//		base_type: base type of the Token
	//		token_val: the actual value of the token
	//return:
	//		Object of from Token class
	token_type := get_child_attribute(base_type, token_val)
	return &Token{token_base_type: base_type, token_value: token_val, token_type: token_type}
}
