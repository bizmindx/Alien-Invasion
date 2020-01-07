package types

import (
	"fmt"
)

// Agent has a name, some state and points to a Node
type Agent struct {
	Name  string
	Flags map[string]bool
	Node  *Node
}

// NewAgent creates an Agent with a name
func NewAgent(name string) Agent {
	return Agent{
		Name:  name,
		Flags: make(map[string]bool),
	}
}

// String representation for an Agent
func (a *Agent) String() string {
	return fmt.Sprintf("name=%s node={%s}\n", a.Name, a.Node)
}
