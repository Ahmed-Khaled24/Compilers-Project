package Parser

func (p *Parser) Entry(token *Node) Node {
	return p.StmtSequence(token)
}

func (p *Parser) StmtSequence(token *Node) Node {
	node := NewNode("StmtSequence", "", nil)
	node.AddChild(p.Statement(token))
	for p.CurrentTokenIndex < len(p.CurrentToken.Children) && p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == ";" {
		p.CurrentTokenIndex++
		node.AddChild(p.Statement(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	}
	return node
}

func (p *Parser) Statement(token *Node) Node {
	switch token.NodeValue {
	case "if":
		return p.IfStmt(token)
	case "repeat":
		return p.RepeatStmt(token)
	case "identifier":
		return p.AssignStmt(token)
	case "read":
		return p.ReadStmt(token)
	case "write":
		return p.WriteStmt(token)
	default:
		panic("Unexpected token")
	}
}

func (p *Parser) IfStmt(token *Node) Node {
	node := NewNode("IfStmt", "", nil)
	node.AddChild(*token)
	p.CurrentTokenIndex++
	node.AddChild(p.Exp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	p.CurrentTokenIndex++
	node.AddChild(p.StmtSequence(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	if p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "else" {
		p.CurrentTokenIndex++
		node.AddChild(p.StmtSequence(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	}
	p.CurrentTokenIndex++
	return node
}

func (p *Parser) RepeatStmt(token *Node) Node {
	node := NewNode("RepeatStmt", "", nil)
	node.AddChild(*token)
	p.CurrentTokenIndex++
	node.AddChild(p.StmtSequence(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	p.CurrentTokenIndex++
	node.AddChild(p.Exp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	p.CurrentTokenIndex++
	return node
}

func (p *Parser) AssignStmt(token *Node) Node {
	node := NewNode("AssignStmt", "", nil)
	node.AddChild(*token)
	p.CurrentTokenIndex++
	node.AddChild(p.Exp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	p.CurrentTokenIndex++
	return node
}

func (p *Parser) ReadStmt(token *Node) Node {
	node := NewNode("ReadStmt", "", nil)
	node.AddChild(*token)
	p.CurrentTokenIndex++
	return node
}

func (p *Parser) WriteStmt(token *Node) Node {
	node := NewNode("WriteStmt", "", nil)
	node.AddChild(*token)
	p.CurrentTokenIndex++
	node.AddChild(p.Exp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	p.CurrentTokenIndex++
	return node
}

func (p *Parser) Exp(token *Node) Node {
	node := NewNode("Exp", "", nil)
	node.AddChild(p.SimpleExp(token))
	for p.CurrentTokenIndex < len(p.CurrentToken.Children) && p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "<" {
		node.AddChild(*token)
		p.CurrentTokenIndex++
		node.AddChild(p.SimpleExp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	}
	return node
}

func (p *Parser) SimpleExp(token *Node) Node {
	node := NewNode("SimpleExp", "", nil)
	node.AddChild(p.Term(token))
	for p.CurrentTokenIndex < len(p.CurrentToken.Children) && (p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "+" || p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "-") {
		node.AddChild(*token)
		p.CurrentTokenIndex++
		node.AddChild(p.Term(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	}
	return node
}

func (p *Parser) Term(token *Node) Node {
	node := NewNode("Term", "", nil)
	node.AddChild(p.Factor(token))
	for p.CurrentTokenIndex < len(p.CurrentToken.Children) && (p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "*" || p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "/") {
		node.AddChild(*token)
		p.CurrentTokenIndex++
		node.AddChild(p.Factor(&p.CurrentToken.Children[p.CurrentTokenIndex]))
	}
	return node
}

func (p *Parser) Factor(token *Node) Node {
	node := NewNode("Factor", "", nil)
	if p.CurrentToken.Children[p.CurrentTokenIndex].NodeValue == "(" {
		node.AddChild(*token)
		p.CurrentTokenIndex++
		node.AddChild(p.Exp(&p.CurrentToken.Children[p.CurrentTokenIndex]))
		p.CurrentTokenIndex++
	} else {
		node.AddChild(*token)
		p.CurrentTokenIndex++
	}
	return node
}
