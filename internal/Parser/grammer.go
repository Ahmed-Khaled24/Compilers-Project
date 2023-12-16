package Parser

import "github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"

// define the parser of the tiny language
func (p *Parser) Program() *Node {
	node := *p.StmtSequence()
	return &node
}

func (p *Parser) StmtSequence() *Node {
	node := p.Statement()
	entry := node

	for p.CurrentToken().TokenType == "SEMICOLON" {
		p.Match("SEMICOLON")
		node.Next = p.Statement()
		node = node.Next
	}
	return entry
}

func (p *Parser) Statement() *Node {
	node := NewNode("Statement", "Statement")
	switch p.CurrentToken().TokenType {
	case "IF":
		node = *p.IfStatement()
	case "REPEAT":
		node = *p.RepeatStatement()
	case "IDENTIFIER":
		node = *p.AssignStatement()
	case "READ":
		node = *p.ReadStatement()
	case "WRITE":
		node = *p.WriteStatement()
	}
	return &node
}

func (p *Parser) IfStatement() *Node {
	node := NewNode("if", "if")
	p.Match("IF")
	node.AddChild(p.Exp())
	p.Match("THEN")
	node.AddChild(p.StmtSequence())
	if p.CurrentToken().TokenType == "ELSE" {
		node.AddChild(p.Match("ELSE"))
		p.StmtSequence()
	}
	p.Match("END")
	return &node
}

func (p *Parser) RepeatStatement() *Node {
	node := NewNode("RepeatStatement", "repeat")
	p.Match("REPEAT")
	node.AddChild(p.StmtSequence())
	p.Match("UNTIL")
	node.AddChild(p.Exp())
	return &node
}

func (p *Parser) AssignStatement() *Node {
	node := NewNode("assign", "assign")
	node.NodeValue = p.Match("IDENTIFIER").NodeValue
	p.Match("ASSIGNMENT")
	node.AddChild(p.Exp())
	return &node
}

func (p *Parser) ReadStatement() *Node {
	node := NewNode("read", "read")
	p.Match("READ")
	node.NodeValue = p.CurrentToken().TokenValue
	p.Match("IDENTIFIER")
	return &node
}

func (p *Parser) WriteStatement() *Node {
	node := NewNode("write", "write")
	p.Match("WRITE")
	node.AddChild(p.Exp())
	return &node
}

func (p *Parser) Exp() *Node {
	node := NewNode("Exp", "Exp")
	child := *p.SimpleExp()
	if p.CurrentToken().TokenType == "LESSTHAN" || p.CurrentToken().TokenType == "EQUAL" {
		node = *p.ComparisonOp()
		node.AddChild(&child)
		node.AddChild(p.SimpleExp())
	} else {
		node = child
	}

	return &node
}

func (p *Parser) ComparisonOp() *Node {
	node := NewNode("ComparisonOp", "ComparisonOp")
	switch p.CurrentToken().TokenType {
	case "LESSTHAN":
		node.NodeValue = "<"
		p.Match("LESSTHAN")
	case "EQUAL":
		node.NodeValue = "="
		p.Match("EQUAL")
	}
	return &node
}

func (p *Parser) SimpleExp() *Node {
	factor := *p.Term()
	node := factor
	for p.CurrentToken().TokenType == "PLUS" || p.CurrentToken().TokenType == "MINUS" {
		node = *p.AddOp()
		node.AddChild(&factor)
		node.AddChild(p.Term())
	}
	return &node
}

func (p *Parser) AddOp() *Node {
	node := NewNode("AddOp", "AddOp")
	switch p.CurrentToken().TokenType {
	case "PLUS":
		node.NodeValue = "+"
		p.Match("PLUS")
	case "MINUS":
		node.NodeValue = "-"
		p.Match("MINUS")
	}
	return &node
}

func (p *Parser) Term() *Node {
	factor := *p.Factor()
	node := factor
	for p.CurrentToken().TokenType == "MULT" || p.CurrentToken().TokenType == "DIV" {
		node = *p.MultOp()
		node.AddChild(&factor)
		node.AddChild(p.Factor())
	}
	return &node
}

func (p *Parser) MultOp() *Node {
	node := NewNode("MultOp", "MultOp")
	switch p.CurrentToken().TokenType {
	case "MULT":
		node.NodeValue = "*"
		p.Match("MULT")
	case "DIV":
		node.NodeValue = "/"
		p.Match("DIV")
	}
	return &node
}

func (p *Parser) Factor() *Node {
	node := NewNode("Factor", "Factor")
	switch p.CurrentToken().TokenType {
	case "OPENBRACKET":
		node.AddChild(p.Match("OPENBRACKET"))
		node.AddChild(p.Exp())
		node.AddChild(p.Match("CLOSEDBRACKET"))
	case "NUMBER":
		node = *p.Match("NUMBER")
	case "IDENTIFIER":
		node = *p.Match("IDENTIFIER")
	}
	return &node
}

func (p *Parser) Match(tokenType string) *Node {
	currentToken := p.CurrentToken()
	node := NewNode(tokenType, currentToken.TokenValue)
	if currentToken.TokenType == tokenType {
		p.Pointer++
	}
	return &node
}

func (p *Parser) CurrentToken() Scanner.Token {
	if p.Pointer >= len(p.TokenList) {
		return Scanner.Token{}
	}
	return p.TokenList[p.Pointer]
}
