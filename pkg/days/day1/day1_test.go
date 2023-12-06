package day1

import (
	"testing"
)

func Test_searchNumbers(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "two1nine",
			want:  29,
		},
		{
			input: "eightwothree",
			want:  83,
		},
		{
			input: "abcone2threexyz",
			want:  13,
		},
		{
			input: "xtwone3four",
			want:  24,
		},
		{
			input: "4nineeightseven2",
			want:  42,
		},
		{
			input: "zoneight234",
			want:  14,
		},
		{
			input: "7pqrstsixteen",
			want:  76,
		},
	}

	for _, tt := range tests {
		got, err := searchNumbers(tt.input, true)
		if err != nil {
			t.Errorf("searchNumbers(%q) returned an error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Errorf("searchNumbers(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func Test_searchWord(t *testing.T) {
	tests := []struct {
		s         string
		i         int
		reverse   bool
		wantFound bool
		wantVal   rune
	}{
		{
			s:         "two1nine",
			i:         0,
			reverse:   true,
			wantFound: false,
			wantVal:   rune('0'),
		},
		{
			s:         "two1nine",
			i:         2,
			reverse:   true,
			wantFound: true,
			wantVal:   rune('2'),
		},
		{
			s:         "eightwothree",
			i:         4,
			reverse:   true,
			wantFound: true,
			wantVal:   rune('8'),
		},
		{
			s:         "two1nine",
			i:         0,
			reverse:   false,
			wantFound: true,
			wantVal:   rune('2'),
		},
		{
			s:         "eightwothree",
			i:         0,
			reverse:   false,
			wantFound: true,
			wantVal:   rune('8'),
		},
		{
			s:         "two1nine",
			i:         1,
			reverse:   false,
			wantFound: false,
			wantVal:   rune('0'),
		},
		{
			s:         "twi",
			i:         0,
			reverse:   false,
			wantFound: false,
			wantVal:   rune('0'),
		},
	}

	for _, tt := range tests {
		gotFound, gotVal := searchWord(tt.s, tt.i, tt.reverse)
		if gotFound != tt.wantFound || gotVal != tt.wantVal {
			t.Errorf("searchWord(%q, %d, %t) = %t, %d, want %t, %d", tt.s, tt.i, tt.reverse, gotFound, gotVal, tt.wantFound, tt.wantVal)
		}
	}
}
