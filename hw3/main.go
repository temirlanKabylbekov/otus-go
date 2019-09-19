package main

import (
	"math"
	"sort"
	"strings"
)

type Pair struct {
	Key 	string
	Value 	int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value || (p[i].Value == p[j].Value && p[i].Key < p[j].Key) }

func frequentWords(text string, count int) []string {
	var whitespaces int
	for _, rune := range text {
		if string(rune) == " " {
			whitespaces += 1
		}
	}

	counter := make(map[string]int, whitespaces + 1)
	words := strings.Split(text, " ")
	for _, word := range words {
		if word != " " && len(word) != 0 {
			counter[word] += 1
		}
	}

	pairs := make(PairList, len(counter))
	i := 0
	for key, value := range counter {
		pairs[i] = Pair{key, value}
		i += 1
	}

	sort.Sort(sort.Reverse(pairs))

	result := make([]string, int(math.Min(float64(count), float64(len(counter)))))
	for i, pair := range pairs {
		if i < count {
			result[i] = pair.Key
		}
	}
	return result
}
