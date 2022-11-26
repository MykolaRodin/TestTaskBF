package q2

import (
	"testing"
)

func Test_RecursiveAlgorithm(t *testing.T) {
	testCases := []struct {
		given    int
		expected []int
	}{
		{
			given:    9,
			expected: []int{2, 4, 9},
		}, {
			given:    2,
			expected: []int{2},
		}, {
			given:    4,
			expected: []int{2, 4},
		}, {
			given:    18,
			expected: []int{2, 4, 9, 18},
		}, {
			given:    19,
			expected: []int{2, 4, 9, 19},
		},
	}

	for _, testCase := range testCases {
		actual := RecursiveAlgorithm(testCase.given)

		for i := 0; i < len(testCase.expected); i++ {
			if actual[i] != testCase.expected[i] {
				t.Errorf("RecursiveAlgorithm(%d) returned:\n%v;\nwant:\n%v", testCase.given, actual, testCase.expected)
			}
		}
	}
}
