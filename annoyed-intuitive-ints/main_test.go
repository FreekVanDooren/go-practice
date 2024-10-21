package main

import (
	"testing"
)

func TestAskForHelp1(t *testing.T) {
	expected := float64(7)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 4); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1a(t *testing.T) {
	expected := float64(5)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1b(t *testing.T) {
	expected := float64(5)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 2); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1c(t *testing.T) {
	expected := float64(4)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 1); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1d(t *testing.T) {
	expected := float64(7)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 5); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1e(t *testing.T) {
	expected := float64(8)
	annoyances := []float64{1, 2, 3, 4}
	deltas := []float64{2, 3, 4, 5}
	nextAnnoyances := []float64{3, 5, 7, 9}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 6); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp2(t *testing.T) {
	expected := float64(1002)
	annoyances := []float64{1, 1000}
	deltas := []float64{1000, 1}
	nextAnnoyances := []float64{1001, 1001}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3(t *testing.T) {
	expected := float64(5)
	annoyances := []float64{1, 2}
	deltas := []float64{1, 2}
	nextAnnoyances := []float64{2, 4}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 5); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3a(t *testing.T) {
	expected := float64(4)
	annoyances := []float64{1, 2}
	deltas := []float64{1, 2}
	nextAnnoyances := []float64{2, 4}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3b(t *testing.T) {
	expected := float64(4)
	annoyances := []float64{1, 2}
	deltas := []float64{1, 2}
	nextAnnoyances := []float64{2, 4}
	if got := askForHelp(annoyances, deltas, nextAnnoyances, 4); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

//func Test_readData(t *testing.T) {
//	type wants struct {
//		coworkers  []*coworker
//		assignment *assignment
//	}
//	tests := []struct {
//		name  string
//		input io.Reader
//		want  wants
//	}{{
//		name: "simple",
//		input: strings.NewReader(`3 2
//1 1000
//1000 1
//`),
//		want: wants{
//			coworkers: []*coworker{{
//				1, 1000,
//			}, {
//				1000, 1,
//			},
//			},
//			assignment: &assignment{
//				3, 2,
//			},
//		},
//	}, {
//		name: "with extra lines",
//		input: strings.NewReader(`
//3 2
//
//1 1000
//
//1000 1
//`),
//		want: wants{
//			coworkers: []*coworker{{
//				1, 1000,
//			}, {
//				1000, 1,
//			},
//			},
//			assignment: &assignment{
//				3, 2,
//			},
//		},
//	},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			a, c := readData(tt.input)
//			if !reflect.DeepEqual(a, tt.want.assignment) {
//				t.Errorf("readData() got = %v, want %v", a, tt.want.assignment)
//			}
//			if !reflect.DeepEqual(c, tt.want.coworkers) {
//				t.Errorf("readData() got = %v, want %v", c, tt.want.coworkers)
//			}
//		})
//	}
//}
//
//func TestIntegration(t *testing.T) {
//	a, c := readData(strings.NewReader(`
//3 2
//
//1 1000
//
//1000 1
//`))
//	wantedAssignment := &assignment{
//		3, 2,
//	}
//	if !reflect.DeepEqual(a, wantedAssignment) {
//		t.Errorf("IT got = %v, want %v", a, wantedAssignment)
//	}
//	wantedCoworkers := []*coworker{{
//		1, 1000,
//	}, {
//		1000, 1,
//	},
//	}
//	if !reflect.DeepEqual(c, wantedCoworkers) {
//		t.Errorf("IT got = %v, want %v", c, wantedCoworkers)
//	}
//
//	annoyance := askForHelp(c, a.helpsNeeded)
//	if annoyance != 1002 {
//		t.Errorf("IT got = %v, want %v", annoyance, 1002)
//	}
//}
