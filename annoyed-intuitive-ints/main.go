package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, annoyances, deltas, nextAnnoyances := readData(os.Stdin)
	totalAnnoyance := askForHelp(annoyances, deltas, nextAnnoyances, a.helpsNeeded)
	fmt.Println(totalAnnoyance)
}

func readData(reader io.Reader) (*assignment, []float64, []float64, []float64) {
	scanner := bufio.NewScanner(reader)
	a := readAssignment(scanner)
	annoyances, deltas, nextAnnoyances := readCoworkers(a, scanner)
	return a, annoyances, deltas, nextAnnoyances
}

func readAssignment(scanner *bufio.Scanner) *assignment {
	for i := 0; i < 10000; i++ {
		scanner.Scan()
		text := scanner.Text()
		if containsData(text) {
			a, err := newAssignment(text)
			if err != nil {
				panic(err)
			}
			return a
		}
	}
	panic(errors.New("Read a 1000 lines, but no assignment"))
}

func readCoworkers(a *assignment, scanner *bufio.Scanner) ([]float64, []float64, []float64) {
	annoyances := make([]float64, 0)
	deltas := make([]float64, 0)
	nextAnnoyances := make([]float64, 0)
	for {
		scanner.Scan()
		text := scanner.Text()

		if containsData(text) {
			annoyance, delta, err := parseLine(text)
			if err != nil {
				panic(err)
			}
			annoyances = append(annoyances, annoyance)
			deltas = append(deltas, delta)
			nextAnnoyances = append(nextAnnoyances, delta+annoyance)

		}
		if len(annoyances) == a.coworkers {
			break
		}
	}
	return annoyances, deltas, nextAnnoyances
}

func containsData(text string) bool {
	return len(strings.TrimSpace(text)) > 0
}

type assignment struct {
	helpsNeeded int
	coworkers   int
}

func (a *assignment) GoString() string {
	return fmt.Sprintf("%+v", *a)
}

func newAssignment(text string) (*assignment, error) {
	helpsNeeded, coworkers, err := parseLine(text)
	if err != nil {
		return nil, err
	}
	return &assignment{
		helpsNeeded: int(helpsNeeded),
		coworkers:   int(coworkers),
	}, nil
}

func parseLine(text string) (float64, float64, error) {
	parts := strings.Split(text, " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected 2, but got %d parts", len(parts))
	}
	nr1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, err
	}
	nr2, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, err
	}
	return nr1, nr2, nil
}

func askForHelp(annoyances, deltas, nextAnnoyances []float64, helpsNeeded int) float64 {
	for i := 0; i < helpsNeeded; i++ {
		leastAnnoyed := nextAnnoyances[0]
		var index int
		for j, nextAnnoyance := range nextAnnoyances[1:] {
			leastAnnoyed = math.Min(nextAnnoyance, leastAnnoyed)
			if leastAnnoyed == nextAnnoyance {
				index = j + 1
			}
		}
		annoyances[index] = nextAnnoyances[index]
		nextAnnoyances[index] += deltas[index]
	}
	var mostAnnoyed float64
	for _, annoyance := range annoyances {
		mostAnnoyed = math.Max(mostAnnoyed, annoyance)
	}
	return mostAnnoyed
}
