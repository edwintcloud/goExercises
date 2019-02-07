package scrabble

import "strings"

// Score computes the scrabble score from a word
func Score(word string) int {
	score := 0
	word = strings.ToLower(word)
	scores := make(map[string]int)
	scores["a"] = 1
	scores["b"] = 3
	scores["c"] = 3
	scores["d"] = 2
	scores["e"] = 1
	scores["f"] = 4
	scores["g"] = 2
	scores["h"] = 4
	scores["i"] = 1
	scores["j"] = 8
	scores["k"] = 5
	scores["l"] = 1
	scores["m"] = 3
	scores["n"] = 1
	scores["o"] = 1
	scores["p"] = 3
	scores["q"] = 10
	scores["r"] = 1
	scores["s"] = 1
	scores["t"] = 1
	scores["u"] = 1
	scores["v"] = 4
	scores["w"] = 4
	scores["x"] = 8
	scores["y"] = 4
	scores["z"] = 10

	// compute score
	for i := 0; i < len(word); i++ {
		score += scores[string(word[i])]
	}

	// return score
	return score
}
