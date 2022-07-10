package leetcode

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	words := strings.Split(sentence, " ")
	// res := []string{}

	for _, v := range dictionary {
		for i, word := range words {
			l := len(v)

			if l <= len(word) && v == word[:l] {
				words[i] = v
			}
		}
	}
	return strings.Join(words, " ")
}

func replaceWords1(dictionary []string, sentence string) string {
	dictionarySet := map[string]bool{}
	for _, s := range dictionary {
		dictionarySet[s] = true
	}

	words := strings.Split(sentence, " ")

	for i, word := range words {
		for j := 1; j < len(word); j++ {
			if dictionarySet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
