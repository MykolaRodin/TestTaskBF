package q1

import "sort"

// sortWords sorts a bunch of words by the number of character “a”s within
// the word (decreasing order). If some words contain the same amount of
// character “a”s then it sorts those words by their lengths.
func SortWords(input []string) []string {
	wordsMap := map[string]int{}

	sort.Slice(input, func(i, j int) bool {
		wordI := input[i]
		wordJ := input[j]

		countI, ok := wordsMap[wordI]
		if !ok {
			countI = countLetterQuantity(wordI, 'a')
			wordsMap[wordI] = countI
		}

		countJ, ok := wordsMap[wordJ]
		if !ok {
			countJ = countLetterQuantity(wordJ, 'a')
			wordsMap[wordJ] = countJ
		}

		if countI > countJ {
			return true
		}
		if countI < countJ {
			return false
		}
		return len(wordI) > len(wordJ)
	})

	return input
}

// countLetterQuantity counts the quantity of letter in the string
func countLetterQuantity(word string, letter byte) int {
	count := 0
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		if word[i] == letter {
			count++
		}
	}

	return count
}
