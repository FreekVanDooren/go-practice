package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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
	arrive             uint
	leave              uint
	stay               uint
	next               *kitten
	hasBed             bool
	kittensCanShareBed uint
	hasPrevious        bool
}

func (c *kitten) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *kitten) nrOfBedsOccupied() uint {
	if c.next != nil {
		return c.next.nrOfBedsOccupied() + 1
	}
	return 1
}

func (c *kitten) setNext(next *kitten) {
	c.next = next
	if next != nil {
		next.hasPrevious = true
	}
}

func (c *kitten) markHasBed() {
	c.hasBed = true
	if c.next != nil {
		c.next.markHasBed()
	}
}

func newKitten(text string) (*kitten, error) {
	arrive, leave, err := parseLine(text)
	if err != nil {
		return nil, err
	}
	return &kitten{
		arrive: arrive,
		leave:  leave,
		stay:   leave - arrive,
	}, nil
}

func parseLine(text string) (uint, uint, error) {
	parts := strings.Split(text, " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected 2, but got %d parts", len(parts))
	}
	nr1, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	nr2, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return uint(nr1), uint(nr2), nil
}

func maximumHoused(kittens []*kitten, beds int) uint {
	arrivals := map[uint][]*kitten{}
	var lastDay uint
	var housed uint
	for _, k := range kittens {
		arrivals[k.arrive] = append(arrivals[k.arrive], k)
		if lastDay < k.leave {
			lastDay = k.leave
		}
	}
	noBeds := kittens
	for b := 0; b < beds; b++ {
		if b != 0 {
			noBeds = []*kitten{}
			arrivals = map[uint][]*kitten{}
			for _, k := range kittens {
				if !k.hasBed {
					noBeds = append(noBeds, k)
					arrivals[k.arrive] = append(arrivals[k.arrive], k)
				}
			}
		}
		for _, k := range noBeds {
			var candidates []*kitten
			for i := k.leave; i < lastDay; i++ {
				candidates = append(candidates, arrivals[i]...)
			}
			if len(candidates) == 0 {
				continue
			}
			var next *kitten
			for _, candidate := range candidates {
				if next == nil {
					next = candidate
					continue
				}
				if candidate.leave <= next.leave && candidate.stay < next.stay {
					next = candidate
				}
			}
			k.setNext(next)
		}
		var bestCandidate *kitten
		for _, k := range noBeds {
			k.kittensCanShareBed = k.nrOfBedsOccupied()
			if bestCandidate == nil {
				bestCandidate = k
				continue
			}
			if k.kittensCanShareBed > bestCandidate.kittensCanShareBed {
				bestCandidate = k
			}
		}
		if bestCandidate != nil {
			housed += bestCandidate.kittensCanShareBed
			bestCandidate.markHasBed()
		}
	}
	return housed
}
