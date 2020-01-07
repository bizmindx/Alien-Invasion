package simulation

import (
	"fmt"
	"math/rand"
	"sort"
	"alien-invasion/simulation/types"
	"alien-invasion/util"
)

type (
	// Alien attacking City
	Alien = types.Alien
	// City connected to other Cities
	City = types.City
	// World map of Cities
	World = types.World
)

// Aliens is a collection of all Aliens
type Aliens []*Alien

// AlienName maps all Aliens by name
type AlienName map[string]*Alien

// CityDefense maps Aliens by City
type CityDefense map[string]AlienName

// Simulation struct represents a running invasion game
type Simulation struct {
	// Simulation config
	R            *rand.Rand
	Iteration    int
	EndIteration int
	// World state
	World
	Aliens
	CityDefense
}

// NoOpReason why Alien did not make a move
type NoOpReason uint8

// NoOpError when next move can not be made
type NoOpError struct {
	reason NoOpReason
}

const (
	// NoOpAlienDead when Alien is Dead
	NoOpAlienDead NoOpReason = iota
	// NoOpAlienTrapped when Alien is Trapped
	NoOpAlienTrapped
	// NoOpWorldDestroyed when World is destroyed
	NoOpWorldDestroyed
	// NoOpMessage when no-op
	NoOpMessage = " || NO move! %s\n"
)

// Error string representation
func (err *NoOpError) Error() string {
	return fmt.Sprintf("Simulator no-op with reason: %d", err.reason)
}

// NewSimulation inits a new Simulation instance
func NewSimulation(r *rand.Rand, endIteration int, world World, aliens Aliens) Simulation {
	return Simulation{
		R:            r,
		Iteration:    0,
		EndIteration: endIteration,
		World:        world,
		Aliens:       aliens,
		CityDefense:  make(CityDefense),
	}
}

// Start the alien invasion
func (s *Simulation) Start() error {
	for ; s.Iteration < s.EndIteration; s.Iteration++ {
		// Shuffle cards every iteration
		picks := util.ShuffleLen(len(s.Aliens), s.R)
		// Aliens make their moves
		noOpRound := true
		for _, p := range picks {
			if err := s.MoveAlien(s.Aliens[p]); err != nil {
				if _, ok := err.(*NoOpError); ok {
					// Alien made no move
					continue
				}
				return err
			}
			// If just one move is made, we continue the simulation
			noOpRound = false
		}
		// Check if last iteration was empty (no moves)
		if noOpRound {
			return nil
		}
	}
	// Game Over
	return nil
}

// MoveAlien moves the Alien position in the simulation
func (s *Simulation) MoveAlien(alien *Alien) error {
	from, to, err := s.pickMove(alien)
	if err != nil {
		return err
	}
	// Move
	alien.InvadeCity(to)
	if from != nil {
		// Move from City
		delete(s.CityDefense[from.Name], alien.Name)
	}
	// Init city defense
	if s.CityDefense[to.Name] == nil {
		s.CityDefense[to.Name] = make(AlienName)
	}
	// Move to City
	s.CityDefense[to.Name][alien.Name] = alien
	if len(s.CityDefense[to.Name]) > 1 {
		to.Destroy()
		// Kill Aliens and notify
		out := fmt.Sprintf("City named %s has been destroyed by ", to.Name)
		for _, a := range s.CityDefense[to.Name] {
			out += fmt.Sprintf("alien %s and ", a.Name)
			a.Kill()
		}
		out = out[:len(out)-5] + "!\n"
		fmt.Print(out)
	}
	// Done
	return nil
}

// pickMove returns Alien move from City to City
func (s *Simulation) pickMove(alien *Alien) (*City, *City, error) {
	// Check if dead or trapped
	from := alien.City()
	if err := checkAlien(alien); err != nil {
		return from, nil, err
	}
	// At the beginning
	if from == nil {
		to := s.pickAnyCity()
		if to == nil {
			return from, to, &NoOpError{reason: NoOpWorldDestroyed}
		}
		return from, to, nil
	}
	// Move to next City
	to := s.pickConnectedCity(alien)
	return from, to, nil
}

// checkAlien returns NoOpError if Alien dead or trapped
func checkAlien(alien *Alien) *NoOpError {
	if alien.IsDead() {
		return &NoOpError{NoOpAlienDead}
	}
	if alien.IsTrapped() {
		return &NoOpError{NoOpAlienTrapped}
	}
	return nil
}

// pickConnectedCity picks a random road to undestroyed City
func (s *Simulation) pickConnectedCity(alien *Alien) *City {
	// Nil if still not invading
	if !alien.IsInvading() {
		return nil
	}
	// Shuffle roads every pick
	picks := util.ShuffleLen(len(alien.City().Links), s.R)
	// Any undestroyed connected city
	for _, p := range picks {
		key := alien.City().Links[p].Key
		n := alien.City().Nodes[key]
		if c := s.World[n.Name]; !c.IsDestroyed() {
			return c
		}
	}
	// No connected undestroyed City
	return nil
}

// pickAnyCity picks any undestroyed City in the World
func (s *Simulation) pickAnyCity() *City {
	// Any undestroyed city, pick deterministically
	var keys []string
	for k := range s.World {
		if c := s.World[k]; !c.IsDestroyed() {
			keys = append(keys, k)
		}
	}
	// If all Cities destroyed
	if len(keys) == 0 {
		return nil
	}
	// Sort keys for a deterministic pick
	sort.Strings(keys)
	pick := s.R.Intn(len(keys))
	return s.World[keys[pick]]
}
