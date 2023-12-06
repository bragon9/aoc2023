package day2

import "testing"

func Test_checkGame(t *testing.T) {
	games := parseGames([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})

	bagContents := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	tests := []struct {
		game Game
		want bool
	}{
		{games[0], true},
		{games[1], true},
		{games[2], false},
		{games[3], false},
		{games[4], true},
	}

	for _, tt := range tests {
		if got := checkGame(tt.game, bagContents); got != tt.want {
			t.Errorf("checkGame() = %v, want %v", got, tt.want)
		}
	}
}
