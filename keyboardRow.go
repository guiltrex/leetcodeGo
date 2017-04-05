package Leetcode

import "strings"

//verbose solution, requiring further refactor in the future.
func findWords(words []string) []string {
	invMap := map[int][]string{
		1: []string{"z", "x", "c", "v", "b", "n", "m"},
		2: []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
		3: []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	}
	strLineMap := reverseMap(invMap)
	res := []string{}

	for _, word := range words {
		lWord := strings.ToLower(word)
		line := strLineMap[string(lWord[0])]
		valid := true
		for _, ch := range lWord {
			if line != strLineMap[string(ch)] {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, word)
		}
	}
	return res

}

func reverseMap(posChMap map[int][]string) map[string]int {
	res := make(map[string]int)
	for pos, strs := range posChMap {
		for _, str := range strs {
			res[str] = pos
		}
	}
	return res
}
