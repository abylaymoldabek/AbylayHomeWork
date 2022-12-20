package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regexpForWords = regexp.MustCompile(`(\p{L}+-\p{L}+)|(\p{L}+)`)

func Top10(strWords string) []string {
	if strWords == "" {
		return make([]string, 0)
	}
	text := strings.ToLower(strWords)
	words := regexpForWords.FindAllString(text, -1)
	cntMap := make(map[string]int)
	for _, v := range words {
		cntMap[v]++
	}
	keys := make([]string, 0, 10)
	for k := range cntMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if cntMap[keys[i]] == cntMap[keys[j]] {
			return keys[i] < keys[j]
		}
		return cntMap[keys[i]] > cntMap[keys[j]]
	})
	if len(keys) > 10 {
		return keys[:10]
	}
	return keys
}
