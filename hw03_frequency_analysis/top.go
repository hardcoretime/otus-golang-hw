package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	words := strings.Fields(s)
	wordCount := make(map[string]int, len(words))
	for _, word := range words {
		wordCount[word]++
	}

	wordList := make([]Word, len(wordCount))
	wordIndex := 0
	for key, value := range wordCount {
		wordList[wordIndex] = Word{Word: key, Count: value}
		wordIndex++
	}

	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].Count == wordList[j].Count {
			return wordList[i].Word < wordList[j].Word
		}

		return wordList[i].Count > wordList[j].Count
	})

	resultCap := 10
	if len(wordList) < 10 {
		resultCap = len(wordList)
	}
	result := make([]string, 0, resultCap)
	for i := 0; i < resultCap; i++ {
		result = append(result, wordList[i].Word)
	}

	return result
}

type Word struct {
	Word  string
	Count int
}
