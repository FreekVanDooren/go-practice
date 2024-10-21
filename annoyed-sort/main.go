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
	a, coworkers := readData(os.Stdin)
	totalAnnoyance := askForHelp(coworkers, a.helpsNeeded)
	fmt.Println(totalAnnoyance)
}

func readData(reader io.Reader) (*assignment, []*coworker) {
	scanner := bufio.NewScanner(reader)
	a := readAssignment(scanner)
	coworkers := readCoworkers(a, scanner)
	return a, coworkers
}

func readAssignment(scanner *bufio.Scanner) *assignment {
	for i := 0; i < 10000; i++ {
		scanner.Scan()
		text := scanner.Text()
		if containsData(text) {
			return onError(newAssignment(text))
		}
	}
	panic(errors.New("read 1000 lines, but no assignment"))
}

func readCoworkers(a *assignment, scanner *bufio.Scanner) []*coworker {
	coworkers := make([]*coworker, 0)
	for {
		scanner.Scan()
		text := scanner.Text()

		if containsData(text) {
			co := onError(newCoworker(text))
			coworkers = append(coworkers, co)
		}
		if len(coworkers) == a.coworkers {
			break
		}
	}
	return coworkers
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

type assignment struct {
	helpsNeeded int
	coworkers   int
}

func (a *assignment) String() string {
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

type coworker struct {
	annoyance uint64
	delta     uint64
}

func (c *coworker) String() string {
	return fmt.Sprintf("%+v", *c)
}

func newCoworker(text string) (*coworker, error) {
	initial, delta, err := parseLine(text)
	if err != nil {
		return nil, err
	}
	return &coworker{
		annoyance: initial,
		delta:     delta,
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

type nextLeastAnnoyed []*coworker

func (a nextLeastAnnoyed) Len() int { return len(a) }
func (a nextLeastAnnoyed) Less(i, j int) bool {
	return a[i].annoyance+a[i].delta < a[j].annoyance+a[j].delta
}
func (a nextLeastAnnoyed) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type mostAnnoyed []*coworker

func (a mostAnnoyed) Len() int { return len(a) }
func (a mostAnnoyed) Less(i, j int) bool {
	return a[i].annoyance > a[j].annoyance
}
func (a mostAnnoyed) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func askForHelp(coworkers []*coworker, helpsNeeded int) uint64 {
	leastAnnoyed := nextLeastAnnoyed(coworkers)
	for i := 0; i < helpsNeeded; i++ {
		sort.Sort(leastAnnoyed)
		leastAnnoyed := leastAnnoyed[0]
		leastAnnoyed.annoyance += leastAnnoyed.delta
	}
	sort.Sort(mostAnnoyed(coworkers))
	return coworkers[0].annoyance
}