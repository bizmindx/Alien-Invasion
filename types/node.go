package types

import (
	"fmt"
)

// Node has a name and is connected to other Nodes via Links
type Node struct {
	Name  string
	Flags map[string]bool
	Links []*Link
	Nodes map[string]*Node
}

// NewNode creates a Node with a name
func NewNode(name string) Node {
	return Node{
		Name:  name,
		Flags: make(map[string]bool),
		Links: make([]*Link, 0),
		Nodes: make(map[string]*Node),
	}
}

// Connect to Node
func (n *Node) Connect(other *Node) *Link {
	link := NewLink(n.Name, other.Name)
	return n.ConnectVia(&link, other)
}

// ConnectVia via link to Node
func (n *Node) ConnectVia(link *Link, other *Node) *Link {
	if n.Nodes[link.Key] == nil {
		n.Links = append(n.Links, link)
		n.Nodes[link.Key] = other
	}
	return link
}

// String representation for a Node
func (n *Node) String() string {
	var links string
	for k, n := range n.Nodes {
		links += fmt.Sprintf("%s:%s ", k, n.Name)
	}
	if len(n.Nodes) > 0 {
		links = links[:len(links)-1]
	}
	return fmt.Sprintf("name=%s links=map[%s]", n.Name, links)
}
