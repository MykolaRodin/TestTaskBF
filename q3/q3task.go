package q3

import "sort"

// GetMostRepeatedString takes one parameter as slice of strings ans returns
// the most repeated string within a given slice.
// If there are more then one most repeated string (for example,{"red", "pie"}),
// the function returns the first one in the alphabetical order ("pie").
func GetMostRepeatedString(input []string) string {
	inputLen := len(input)
	if inputLen == 0 {
		return ""
	}

	sort.Strings(input)

	maxRepeatedCount := 1
	maxRepeatedString := input[0]
	prevRepeatedCount := 1
	prevString := input[0]

	for i := 1; i < inputLen; i++ {
		currString := input[i]
		if prevString != currString {
			if maxRepeatedCount < prevRepeatedCount {
				maxRepeatedCount = prevRepeatedCount
				maxRepeatedString = prevString
			}
			prevRepeatedCount = 1
			prevString = currString
		} else {
			prevRepeatedCount++

			if i == inputLen-1 && maxRepeatedCount < prevRepeatedCount {
				maxRepeatedCount = prevRepeatedCount
				maxRepeatedString = prevString
			}
		}
	}

	return maxRepeatedString
}
