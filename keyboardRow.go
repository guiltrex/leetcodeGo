package Leetcode

import "strings"

//verbose solution, requiring further refactor in the future.
func findWords(words []string) []string {
	posMatrix := [][]string{
		[]string{"z", "x", "c", "v", "b", "n", "m"},
		[]string{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
		[]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	}
	strLineMap := buildMap(posMatrix)
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

func buildMap(posMatrix [][]string) map[string]int {
	res := make(map[string]int)
	for pos, strs := range posMatrix {
		for _, str := range strs {
			res[str] = pos
		}
	}
	return res
}
