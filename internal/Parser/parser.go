package Parser

import (
	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)

type Parser struct {
	Pointer   int
	TreeEntry *Node
	TokenList []Scanner.Token
}

func NewParser(tokenList []Scanner.Token) *Parser {
	return &Parser{
		Pointer:   0,
		TokenList: tokenList,
	}
}

func (p *Parser) Parse() *Node {
	p.TreeEntry = p.Program()
	return p.TreeEntry
}
