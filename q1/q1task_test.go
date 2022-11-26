package q1

import "testing"

func Test_SortWords(t *testing.T) {
	testCases := []struct {
		given    []string
		expected []string
	}{
		{
			given:    []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			expected: []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"},
		}, {
			given:    []string{"aab", "aba", "aaa"},
			expected: []string{"aaa", "aab", "aba"},
		}, {
			given:    []string{"b"},
			expected: []string{"b"},
		}, {
			given:    []string{"a"},
			expected: []string{"a"},
		}, {
			given:    []string{""},
			expected: []string{""},
		}, {
			given:    []string{},
			expected: []string{},
		},
	}

	for _, testCase := range testCases {
		actual := SortWords(testCase.given)

		for i := 0; i < len(testCase.given); i++ {
			if actual[i] != testCase.expected[i] {
				t.Errorf("sortWords(%v) returned:\n%v;\nwant:\n%v", testCase.given, actual, testCase.expected)
			}
		}
	}
}
