package main

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

//go:embed data/*.ans
//go:embed data/*.in
var testData embed.FS

func TestCases(t *testing.T) {
	var testFiles []string
	err := fs.WalkDir(testData, "data", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		testFiles = append(testFiles, path)
		return nil
	})
	if err != nil {
		t.Errorf("No error expected, got %v", err)
	}
	testCases := map[string]string{}
	answer := func(fileName string) string {
		answerFile := strings.ReplaceAll(fileName, ".in", ".ans")
		for _, testFile := range testFiles {
			if testFile == answerFile {
				return testFile
			}
		}
		return ""
	}
	for _, fileName := range testFiles {
		if strings.HasSuffix(fileName, ".in") {
			testCases[fileName] = answer(fileName)
		}
	}
	for caseFile, answerFile := range testCases {
		testCase, err := testData.ReadFile(caseFile)
		if err != nil {
			t.Errorf("No error expected, got %v", err)
		}
		wantRaw, err := testData.ReadFile(answerFile)
		if err != nil {
			t.Errorf("No error expected, got %v", err)
		}
		wantStr := strings.TrimSpace(string(wantRaw))
		want, err := strconv.ParseUint(wantStr, 10, 32)
		if err != nil {
			t.Errorf("No error expected, got %v", err)
		}
		t.Run(caseFile, func(t *testing.T) {
			a, coworkers := readData(bytes.NewReader(testCase))
			got := askForHelp(coworkers, a.helpsNeeded)
			if uint(want) != got {
				t.Errorf("Want %v, got %v", want, got)
			}
		})
	}
}

func TestAskForHelp1(t *testing.T) {
	expected := uint(7)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 4); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1a(t *testing.T) {
	expected := uint(5)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1b(t *testing.T) {
	expected := uint(5)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 2); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1c(t *testing.T) {
	expected := uint(4)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 1); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1d(t *testing.T) {
	expected := uint(7)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 5); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp1e(t *testing.T) {
	expected := uint(8)
	input := []*coworker{{
		1, 2, 3,
	}, {
		2, 3, 5,
	}, {
		3, 4, 7,
	}, {
		4, 5, 8,
	}}
	if got := askForHelp(input, 6); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func TestAskForHelp2(t *testing.T) {
	expected := uint(1002)
	input := []*coworker{{
		1, 1000, 1001,
	}, {
		1000, 1, 1001,
	}}
	if got := askForHelp(input, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3(t *testing.T) {
	expected := uint(5)
	input := []*coworker{{
		1, 1, 2,
	}, {
		2, 2, 4,
	}}
	if got := askForHelp(input, 5); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3a(t *testing.T) {
	expected := uint(4)
	input := []*coworker{{
		1, 1, 2,
	}, {
		2, 2, 4,
	}}
	if got := askForHelp(input, 3); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}
func TestAskForHelp3b(t *testing.T) {
	expected := uint(4)
	input := []*coworker{{
		1, 1, 2,
	}, {
		2, 2, 4,
	}}
	if got := askForHelp(input, 4); got != expected {
		t.Errorf("want %v, got %v", expected, got)
	}
}

func Test_readData(t *testing.T) {
	type wants struct {
		coworkers  []*coworker
		assignment *assignment
	}
	tests := []struct {
		name  string
		input io.Reader
		want  wants
	}{{
		name: "simple",
		input: strings.NewReader(`3 2
1 1000
1000 1
`),
		want: wants{
			coworkers: []*coworker{{
				1, 1000, 1001,
			}, {
				1000, 1, 1001,
			},
			},
			assignment: &assignment{
				3, 2,
			},
		},
	}, {
		name: "with extra lines",
		input: strings.NewReader(`
3 2

1 1000

1000 1
`),
		want: wants{
			coworkers: []*coworker{{
				1, 1000, 1001,
			}, {
				1000, 1, 1001,
			},
			},
			assignment: &assignment{
				3, 2,
			},
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, c := readData(tt.input)
			if !reflect.DeepEqual(a, tt.want.assignment) {
				t.Errorf("readData() got = %v, want %v", a, tt.want.assignment)
			}
			if !reflect.DeepEqual(c, tt.want.coworkers) {
				t.Errorf("readData() got = %v, want %v", c, tt.want.coworkers)
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	a, c := readData(strings.NewReader(`
3 2

1 1000

1000 1
`))
	wantedAssignment := &assignment{
		3, 2,
	}
	if !reflect.DeepEqual(a, wantedAssignment) {
		t.Errorf("IT got = %v, want %v", a, wantedAssignment)
	}
	wantedCoworkers := []*coworker{{
		1, 1000, 1001,
	}, {
		1000, 1, 1001,
	},
	}
	if !reflect.DeepEqual(c, wantedCoworkers) {
		t.Errorf("IT got = %v, want %v", c, wantedCoworkers)
	}

	annoyance := askForHelp(c, a.helpsNeeded)
	if annoyance != 1002 {
		t.Errorf("IT got = %v, want %v", annoyance, 1002)
	}
}
