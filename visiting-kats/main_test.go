package main

import (
	"strings"
	"testing"
)

func TestKittens(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  uint64
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
