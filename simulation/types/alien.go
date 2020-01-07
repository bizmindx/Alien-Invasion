package types

import (
	"fmt"

	"alien-invasion/types"
)

const (
	// FlagDead is a flag used to mark dead Aliens
	FlagDead string = "dead"
)

// Alien can be dead or alive and occupying a City
type Alien struct {
	types.Agent
	city *City
}

// NewAlien creates an Alien with a name and default flags
func NewAlien(name string) Alien {
	// Flags FlagDead default is false
	return Alien{
		Agent: types.NewAgent(name),
	}
}

// InvadeCity change City this Alien is occupying
func (a *Alien) InvadeCity(city *City) {
	a.Node = &city.Node
	a.city = city
}

// City returns City this Alien is occupying
func (a *Alien) City() *City {
	return a.city
}

// IsDead checks if Alien died
func (a *Alien) IsDead() bool {
	return a.Flags[FlagDead]
}

// Kill Alien makes it dead
func (a *Alien) Kill() {
	a.Flags[FlagDead] = true
}

// IsInvading checks if Alien is currently invading a City
func (a *Alien) IsInvading() bool {
	return a.Node != nil
}

// IsTrapped checks if Alien is trapped in a City with no roads out
func (a *Alien) IsTrapped() bool {
	if !a.IsInvading() {
		return false
	}
	for _, n := range a.City().Nodes {
		c := City{Node: *n}
		if !c.IsDestroyed() {
			return false
		}
	}
	return true
}

// String representation for an Alien
func (a *Alien) String() string {
	return fmt.Sprintf("name=%s city={%s}\n", a.Name, a.city)
}
