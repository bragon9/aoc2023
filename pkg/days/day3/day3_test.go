package day3

import "testing"

func Test_checkNumber(t *testing.T) {
	tests := []struct {
		name              string
		coords            [2]string
		symbolLocationMap map[string]any
		want              bool
	}{
		{
			name:              "Top Left Corner",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"0,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Bottom Left Corner",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"0,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Above",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"2,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Top Right Corner",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"4,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Bottom Right Corner",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"4,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Below",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"2,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Left",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"0,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Right",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{"4,1": struct{}{}},
			want:              true,
		},
		{
			name:              "Top Left Corner 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"0,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Bottom Left Corner 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"0,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Above 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"1,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Top Right Corner 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"2,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Bottom Right Corner 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"2,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Below 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"1,2": struct{}{}},
			want:              true,
		},
		{
			name:              "Left 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"0,0": struct{}{}},
			want:              true,
		},
		{
			name:              "Right 1 length",
			coords:            [2]string{"1,1", "1,1"},
			symbolLocationMap: map[string]any{"2,1": struct{}{}},
			want:              true,
		},
		{
			name:              "None",
			coords:            [2]string{"1,1", "3,1"},
			symbolLocationMap: map[string]any{},
			want:              false,
		},
	}

	for _, tt := range tests {
		if got := checkNumber(tt.coords, tt.symbolLocationMap); got != tt.want {
			t.Errorf("checkNumber(%s) = %v, want %v", tt.coords, got, tt.want)
		}
	}
}
