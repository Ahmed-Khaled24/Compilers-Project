package Parser

// define the struct of node
type Node struct {
	NodeType  string
	NodeValue string
	Children  []Node
	Next      *Node
}

func NewNode(nodeType string, nodeValue string) Node {
	return Node{
		NodeType:  nodeType,
		NodeValue: nodeValue,
		Children:  []Node{},
	}
}

func (n *Node) AddChild(child *Node) *Node {
	x := append(n.Children, *child)
	n.Children = x
	return n
}

func (n *Node) AddNext(next *Node) {
	n.Next = next
}
