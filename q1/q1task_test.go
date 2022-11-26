package q1

import "testing"

func Test_SortWords(t *testing.T) {
	inputs := [][]string{
		{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
		{"aab", "aba", "aaa"},
		{"b"},
		{"a"},
		{""},
		{},
	}
	outputs := [][]string{
		{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"},
		{"aaa", "aab", "aba"},
		{"b"},
		{"a"},
		{""},
		{},
	}

	testPosLen := len(inputs)
	for testPos := 0; testPos < testPosLen; testPos++ {
		given := inputs[testPos]
		actual := SortWords(given)
		expected := outputs[testPos]

		givenLen := len(given)
		for i := 0; i < givenLen; i++ {
			if actual[i] != expected[i] {
				t.Errorf("sortWords(%v) returned:\n%v;\nwant:\n%v", given, actual, expected)
			}
		}
	}
}
