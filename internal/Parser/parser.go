package Parser

import (
	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)

type Parser struct {
	Pointer   int
	TreeEntry *Node
	TokenList []Scanner.Token
}


func (p *Parser) Parse(inputString string) *Node {
	tokenList := Scanner.Fsm.Scan(inputString)
	p.TokenList = tokenList
	p.Pointer = 0
	p.TreeEntry = p.Program()
	return p.TreeEntry
}

var ParserSingleton = Parser{}