package Parser

type Parser struct {
	CurrentTokenIndex int
	CurrentToken      *Node
}

func NewParser(entry *Node) Parser {
	return Parser{
		CurrentTokenIndex: 0,
		CurrentToken:      entry,
	}
}

func Parse(tokenList []Node) Node {
	parser := NewParser(&tokenList[0])
	entry := parser.Entry(parser.CurrentToken)
	return entry
}

// define a function to print the tree
func (p *Parser) PrintTree(node *Node, level int) {
	for i := 0; i < level; i++ {
		print("  ")
	}
	println(node.NodeType)
	for _, child := range node.Children {
		p.PrintTree(&child, level+1)
	}
}


// define a function to test the parser and put some test case yourself in an array and then make the test case you want to test uncomment
func (p *Parser) TestPrint() {
	// test case 1
	tokenList := []Node{
		{
			NodeType:  "StmtSequence",
			NodeValue: "",
			Children: []Node{
				{
					NodeType:  "IfStmt",
					NodeValue: "if",
					Children: []Node{
						{
							NodeType:  "Exp",
							NodeValue: "0",
							Children: []Node{
								{
									NodeType:  "SimpleExp",
									NodeValue: "0",
									Children: []Node{
										{
											NodeType:  "Term",
											NodeValue: "0",
											Children: []Node{
												{
													NodeType:  "Factor",
													NodeValue: "0",
													Children: []Node{
														{
															NodeType:  "number",
															NodeValue: "0",
															Children:  []Node{},
															Next:      nil,
														},
													},
													Next: nil,
												},
											},
											Next: nil,
										},
									},
									Next: nil,
								},
							},
							Next: nil,
						},
						{
							NodeType:  "StmtSequence",
							NodeValue: "",
							Children: []Node{
								{
									NodeType:  "AssignStmt",
									NodeValue: "identifier",
									Children: []Node{
										{
											NodeType:  "Exp",
											NodeValue: "1",
											Children: []Node{
												{
													NodeType:  "SimpleExp",
													NodeValue: "1",
													Children: []Node{
														{
															NodeType:  "Term",
															NodeValue: "1",
															Children: []Node{
																{
																	NodeType:  "Factor",
																	NodeValue: "1",
																	Children: []Node{
																		{
																			NodeType:  "number",
																			NodeValue: "1",
																			Children:  []Node{},
																			Next:      nil,
																		},
																	},
																	Next: nil,
																},
															},
															Next: nil,
														},
													},
													Next: nil,
												},
											},
											Next: nil,
										},
									},
									Next: nil,
								},
							},
							Next: nil,
						},
					},
					Next: nil,
				},
			},
			Next: nil,
		},
	}
	// write the Tiny code of tokenlist above in a comment
	// { Sample program in TINY language – computes factorial}
	// read 3;   {input an integer }
	// if  0 < x   then     {  don’t compute if x <= 0 }
	//    fact  := 1;
	//    repeat
	//       fact  := fact *  x;
	//        x  := x  -  1
	//    until {rerer} x  =  0;
	//    write  fact
	// end

	// test the token list above
	p.PrintTree(&tokenList[0], 0)
}
