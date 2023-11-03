package Scanner

import "strings"

type BASETOKENTYPE int64
type ComparsionToken int64
type ArithmeticToken int64
type BracketToken int64
type KeyWords int64

const SEMICOLON string = "SEMICOLON"
const UNKOWN string = "UNKOWN"
const (
	IF KeyWords = iota
	THEN
	ELSE
	END
	REPEAT
	UNTIL
	READ
	WRITE
)
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

func (t KeyWords) String() string {
	switch t {
	case IF:
		return "IF"
	case THEN:
		return "THEN"
	case ELSE:
		return "ELSE"
	case END:
		return "END"
	case REPEAT:
		return "REPEAT"
	case UNTIL:
		return "UNTIL"
	case READ:
		return "READ"
	case WRITE:
		return "WRITE"
	}

	return UNKOWN
}
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

	return UNKOWN
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
	return UNKOWN
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
	return UNKOWN
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
	return UNKOWN
}

func get_child_attribute(base_type BASETOKENTYPE, token_val string) string {
	switch base_type {
	case SPECIALSYMBOL:
		return handle_special_symbols(token_val)
	case RESERVED:
		return handle_revserved_words(token_val)
	case ASSIGNMENT:
		return ASSIGNMENT.String()
	case NUMBER:
		return NUMBER.String()
	case IDENTIFIER:
		return IDENTIFIER.String()
	case ERROR:
		return ERROR.String()
	}
	return UNKOWN
}
func handle_revserved_words(token_val string) string {
		switch token_val {
	case "if":
		return IF.String()
	case "else":
		return ELSE.String()
	case "then":
		return THEN.String()
	case "end":
		return END.String()
	case "repeat":
		return REPEAT.String()
	case "until":
		return UNTIL.String()
	case "read":
		return READ.String()
	case "write":
		return WRITE.String()
	}
	return UNKOWN
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
	case "+":
		return PLUS.String()
	case "-":
		return MINUS.String()
	case "/":
		return DIV.String()
	case "*":
		return MULT.String()
	case "(":
		return OPENBRACKET.String()
	case ")":
		return CLOSEDBRACKET.String()
	case "{":
		return OPENBRACE.String()
	case "}":
		return CLOSEDBRACE.String()
	case ";":
		return SEMICOLON
	}
	return UNKOWN
}

type Token struct {
	TokenBaseType BASETOKENTYPE
	TokenValue    string
	TokenType     string
}

// naming convision: The method or function you want to access should start with an uppercase letter in Go
// constructor for Token struct
// param:
//
//	base_type: base type of the Token
//	token_val: the actual value of the token
//
// return:
//
//	Object of from Token class
func CreateToken(base_type BASETOKENTYPE, token_val string) *Token {
	token_val = strings.TrimSpace(token_val)
	token_type := get_child_attribute(base_type, token_val)
	return &Token{TokenBaseType: base_type, TokenValue: token_val, TokenType: token_type}
}
