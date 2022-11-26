package q3

import "testing"

func Test_GetMostRepeatedString(t *testing.T) {
	testCases := []struct {
		given    []string
		expected string
	}{
		{
			given:    []string{"apple", "pie", "apple", "red", "red", "red"},
			expected: "red",
		},
		{
			given:    []string{"apple", "pie", "apple", "apple", "red", "red", "red"},
			expected: "apple",
		},
		{
			given:    []string{"red", "apple"},
			expected: "apple",
		},
		{
			given:    []string{"red", "pie", "melon"},
			expected: "melon",
		},
	}

	for _, testCase := range testCases {
		actual := GetMostRepeatedString(testCase.given)

		if actual != testCase.expected {
			t.Errorf("GetMostRepeatedString(%v) returned: %s; want: %s", testCase.given, actual, testCase.expected)
		}
	}
}
