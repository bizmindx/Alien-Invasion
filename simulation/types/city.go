package types

import (
	"fmt"

	"alien-invasion/types"
)

const (
	// FlagDestroyed is a flag used to mark destroyed Cities
	FlagDestroyed string = "destroyed"
)

// City has a name and is connected to other Cities via roads
type City struct {
	types.Node
	RoadNames map[string]string
}

// NewCity creates a City with a name and default flags
func NewCity(name string) City {
	// FlagDestroyed default is false
	return City{
		Node:      types.NewNode(name),
		RoadNames: make(map[string]string),
	}
}

// IsDestroyed checks if City is destroyed
func (c *City) IsDestroyed() bool {
	return c.Flags[FlagDestroyed]
}

// Destroy City makes City burn in flames
func (c *City) Destroy() {
	c.Flags[FlagDestroyed] = true
}

// String representation for a City does not print destroyed linked Cities
func (c *City) String() string {
	var links string
	for _, link := range c.Links {
		n := c.Nodes[link.Key]
		other := City{Node: *n}
		// If other City destroyed print nothing
		if other.IsDestroyed() {
			continue
		}
		// If other City survived print Link
		links += fmt.Sprintf("%s=%s ", c.RoadNames[link.Key], other.Name)
	}
	if len(links) == 0 {
		return c.Name
	}
	return fmt.Sprintf("%s %s", c.Name, links[:len(links)-1])
}
