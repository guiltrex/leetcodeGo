package Leetcode

import (
	"sort"
)


//playing with functional programming
func isAnagram1(s string, t string) bool {

	genLess := func(rs []rune) func(int, int) bool {
		return func(i, j int) bool {
			return rs[i] < rs[j]
		}
	}

	sr, tr := []rune(s), []rune(t)
	sort.Slice(sr, genLess(sr))
	sort.Slice(tr, genLess(tr))
	return string(sr) == string(tr)
}

//solution using map
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) { return false }

	smap := make(map[rune]int)
	for _, v := range s {
		smap[v]+=1
	}

	for _, k := range t {
		if v, ok := smap[k]; !ok {
			return false
		} else {
			smap[k] = v-1
		}
	}

	for _, v := range smap{
		if v!= 0 {return false}
	}

	return true
}