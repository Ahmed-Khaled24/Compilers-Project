package Parser

// define the struct of node
type Node struct {
	NodeType  string
	NodeValue string
	Children  *[]Node
	Parent    *Node
	Next      *Node
}

func NewNode(nodeType string, nodeValue string, parent *Node, next *Node) Node {
	return Node{
		NodeType:  nodeType,
		NodeValue: nodeValue,
		Children:  &[]Node{},
		Parent:    parent,
		Next:      next,
	}
}

func (n *Node) AddChild(child Node) {
	*n.Children = append(*n.Children, child)
}
