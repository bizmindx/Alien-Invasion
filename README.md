## Alien Invasion

  

Mad aliens are about to invade the earth and you are tasked with simulating the invasion.

  

### Prerequisite

setup [Golang working environment](https://golang.org/doc/install)
  
  
### usage
Clone the repository and install the dependencies

```bash
git clone https://github.com/codaelux/Alien-Invasion.git
cd Alien-Invasion
```

  

### Run
Getting started
```
$ go run main.go -aliens 20
```
**Options**: View all CLI options
  
```bash
$ go run main.go -help
```
CLI options

    usage of /main:
     
      -aliens int
            number of aliens invading
      -intel string
            a file used to name aliens (default "./files/aliens.txt")
      -movements int
            number of alien movements (default 10000)
      -world string
            a file used as game world map input (default "./files/example.txt")

running with specific alien movement command
``` 
$ go run main.go -aliens 20 -movements 50 
```
running with  alien world path and alien name directory 

    $ go run main.go -aliens 20 -movements 50 -world "./test/example.txt" -intel "./test/alien.txt"

## Testing
Run the test suite with:

```
$ go test ./... -v
