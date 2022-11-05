package words_top

import (
	"sort"
	"strings"
	"unicode"
)

type WordOccurence struct {
	word  string
	count uint
}

func Top10(s string) []string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	countsMap := make(map[string]uint, len(words))

	for _, word := range words {
		lowered_word := strings.ToLower(word)
		countsMap[lowered_word]++
	}

	counts := make([]WordOccurence, 0, len(countsMap))
	for word, count := range countsMap {
		counts = append(counts, WordOccurence{word, count})
	}

	sort.Slice(counts, func(i int, j int) bool {
		a, b := counts[i], counts[j]
		if a.count == b.count {
			return a.word < b.word
		}
		return a.count > b.count
	})

	top_cap := 10
	if top_cap > len(counts) {
		top_cap = len(counts)
	}

	top10 := counts[:top_cap]

	top10strings := make([]string, 0, top_cap)

	for _, word_occurence := range top10 {
		top10strings = append(top10strings, word_occurence.word)
	}

	return top10strings
}
