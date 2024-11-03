package main

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"testing"
	"time"
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
			start := time.Now()
			availability, kittens := readData(bytes.NewReader(testCase))
			fmt.Println(time.Now().Sub(start))
			got := maximumHoused(kittens, availability.beds)
			fmt.Println(time.Now().Sub(start))
			if uint(want) != got {
				t.Errorf("Want %v, got %v", want, got)
			}
		})
	}
}

func TestKittens(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  uint
	}{
		{
			name: "example 1",
			input: `3 1
1 2
2 3
2 3
`,
			want: 2,
		},
		{
			name: "example 2",
			input: `4 1
1 3
4 6
7 8
2 5

`,
			want: 3,
		},
		{
			name: "example 3",
			input: `5 2
1 4
5 9
2 7
3 8
6 10

`,
			want: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			availability, kittens := readData(strings.NewReader(test.input))
			got := maximumHoused(kittens, availability.beds)
			if got != test.want {
				t.Errorf("IT got = %v, want %v", got, test.want)
			}
		})
	}
}
