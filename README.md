# Alien Invasion simulator in Go

This is a _deterministic_ simulator written in Go to simulate an [alien invasion](./CHALLENGE.md).

## Build & Run

To run the `alien-invasion` have a [working Golang environment](https://golang.org/doc/install) installed. If you are all set, just run the following:

```
$ go run main.go
```
This will run the simulation using all defaults and current Unix time as a random source of entropy.

To list all `cli` options ask for help:
```
$ go run main.go -help
Usage of /main:
  -aliens int
        number of aliens invading (default 10)
  -entropy int
        random number used as entropy seed
  -intel string
        a file used to identify aliens (default "./test/aliens.txt")
  -iterations int
        number of iterations (default 10000)
  -simulation string
        name hashed and used as entropy seed
  -world string
        a file used as world map input (default "./test/example.txt")
```

You can run the specific simulation by providing entropy:

```
$ go run main.go -aliens 4 -entropy 123
```

Or provide a simulation name (key) from which entropy will be extracted (sha265):

```
$ go run main.go -aliens 4 -iterations 100 -world "./test/example_2.txt" -simulation "Battle for Cosmos"
```

Reuse the same entropy (or simulation name) to run exactly the same simulation over again. This next command will run the same "Battle for Cosmos" simulation but this time using provided entropy:

```
$ go run main.go -aliens 4 -iterations 100 -world "./test/example_2.txt" -entropy -7645731219066279255
```

## Implementation

The simulation is implemented as a deterministic state machine. While Agents make random decisions on every move, that decision is deterministic and is generated using a pseudorandom number generator. A deterministic application like the one implemented here can be deployed as a blockchain smart contract. Using a platform like [Cosmos](https://cosmos.network/sdk) one could make a deterministic game that solely depends on the provided entropy as an input parameter.

### Random numbers

While [completely random is not really possible](https://www.youtube.com/watch?v=sMb00lz-IfE), we still can have pseudorandom numbers on computers.

The [`math/rand` package](https://golang.org/pkg/math/rand/) provided by the Go Standard Library gives us [pseudo-random number generators (PRNG)](https://en.wikipedia.org/wiki/Pseudorandom_number_generator), also called _deterministic random bit generators_.

As with all pseudo number generators, any number generated through `math/rand` is not really random by default, as being deterministic it will always print the same value each time.

**Notice:** Logs are not deterministic at this point. Nondeterminism is caused by iterating over a hash map in Go.

### State

We use flags map for Agents and Nodes to avoid deleting and removing information from the state graph. This can enable us to implement another simulation with slightly different rules. For example, this other implementation could:

- resurrect Aliens at some point
- allow Aliens to teleport if trapped in a City
- rebuild City when Aliens are gone
- search roads for distant connected undestroyed City to move to (not only next-door-neighbor)

### Next steps

- [ ] More map examples
- [ ] More unit tests
- [ ] Deterministic I/O tests
- [ ] Custom logger and log levels
- [ ] Circle CI tests on commit
- [ ] Codecov test coverage report

## Tests

To run the tests for `alien-invasion` run the following from the root of the repo:

```
$ go test ./... -v
```
