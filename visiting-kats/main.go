package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	availability, kittens := readData(os.Stdin)
	maxHoused := maximumHoused(kittens, availability.beds)
	fmt.Println(maxHoused)
}

func readData(reader io.Reader) (*bedAvailability, []*kitten) {
	scanner := bufio.NewScanner(reader)
	b := readBeds(scanner)
	kittens := readKittens(b, scanner)
	return b, kittens
}

func readBeds(scanner *bufio.Scanner) *bedAvailability {
	for i := 0; i < 10000; i++ {
		scanner.Scan()
		text := scanner.Text()
		if containsData(text) {
			return onError(newBedAvailability(text))
		}
	}
	panic(errors.New("read 1000 lines, but no bedAvailability"))
}

func readKittens(a *bedAvailability, scanner *bufio.Scanner) []*kitten {
	kittens := make([]*kitten, 0)
	for {
		scanner.Scan()
		text := scanner.Text()

		if containsData(text) {
			co := onError(newKitten(text))
			kittens = append(kittens, co)
		}
		if len(kittens) == a.visiting {
			break
		}
	}
	return kittens
}

func containsData(text string) bool {
	return len(strings.TrimSpace(text)) > 0
}

func onError[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

type bedAvailability struct {
	visiting int
	beds     int
}

func (a *bedAvailability) String() string {
	return fmt.Sprintf("%+v", *a)
}

func newBedAvailability(text string) (*bedAvailability, error) {
	visiting, b, err := parseLine(text)
	if err != nil {
		return nil, err
	}
	return &bedAvailability{
		visiting: int(visiting),
		beds:     int(b),
	}, nil
}

type kitten struct {
	arrive uint64
	leave  uint64
	next   *kitten
	hasBed bool
}

func (c *kitten) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *kitten) nrOfBedsOccupied() uint64 {
	c.hasBed = true
	if c.next != nil {
		return c.next.nrOfBedsOccupied() + 1
	}
	return 1
}

func newKitten(text string) (*kitten, error) {
	initial, delta, err := parseLine(text)
	if err != nil {
		return nil, err
	}
	return &kitten{
		arrive: initial,
		leave:  delta,
	}, nil
}

func parseLine(text string) (uint64, uint64, error) {
	parts := strings.Split(text, " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected 2, but got %d parts", len(parts))
	}
	nr1, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	nr2, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return nr1, nr2, nil
}

type byArrival []*kitten

func (a byArrival) Len() int           { return len(a) }
func (a byArrival) Less(i, j int) bool { return a[i].arrive < a[j].arrive }
func (a byArrival) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func maximumHoused(kittens []*kitten, beds int) uint64 {
	sort.Sort(byArrival(kittens))

	for i, k := range kittens {
		k.next = findNext(kittens[i+1:], k)
	}

	c := make(chan uint64)
	for i := 0; i < beds; i++ {
		var nextKittenToBed *kitten
		for _, k := range kittens[i:] {
			if !k.hasBed {
				nextKittenToBed = k
				break
			}
		}
		go func(k *kitten) {
			bedsOccupied := k.nrOfBedsOccupied()
			c <- bedsOccupied
		}(nextKittenToBed)
	}
	var housed uint64
	var valuesReceived int
	for {
		select {
		case bedsOccupied := <-c:
			housed += bedsOccupied
			valuesReceived++
		default:
			if valuesReceived == beds {
				return housed
			}
		}
	}
}

func findNext(kittens []*kitten, current *kitten) *kitten {
	for _, k := range kittens {
		if k.arrive >= current.leave {
			return k
		}
	}
	return nil
}
