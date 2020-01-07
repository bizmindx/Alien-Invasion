package simulation

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestWorldReadFromFile(t *testing.T) {
	testWorldReadFromFile(t, "../test/example.txt", 5)
}

func TestWorldReadFromFile_2(t *testing.T) {
	testWorldReadFromFile(t, "../test/example_2.txt", 6)
}

func TestWorldReadFromFile_3(t *testing.T) {
	testWorldReadFromFile(t, "../test/example_3.txt", 8)
}

func testWorldReadFromFile(t *testing.T, file string, num int) {
	w, input, err := ReadWorldMapFile(file)
	if err != nil {
		t.Errorf("%v: could not read file", err)
		return
	}
	fmt.Printf("Input:\n%s\n\n", input)
	fmt.Printf("World:\n%s\n\n", w)
	if len(w) != num {
		t.Errorf("len(World) = %d; want %d", len(w), num)
		return
	}
}

func TestRandAliens(t *testing.T) {
	testRandAliens(t, 10, 0xffffffff)
}

func TestRandAliensDouble(t *testing.T) {
	testRandAliens(t, 20, 0xffffffff)
}

func testRandAliens(t *testing.T, n int, seed int64) {
	source := rand.NewSource(seed)
	r := rand.New(source)
	aliens := RandAliens(n, r)
	if len(aliens) != n {
		t.Errorf("len(RandAliens(%d, 0xffffffff)) = %d; want %d", n, len(aliens), n)
		return
	}
	fmt.Printf("Aliens:\n%s\n", aliens)
}
