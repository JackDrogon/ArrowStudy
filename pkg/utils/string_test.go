package utils

import "testing"

func Test_AsciiStringReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character string",
			input:    "a",
			expected: "a",
		},
		{
			name:     "multiple character string",
			input:    "hello",
			expected: "olleh",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := asciiStringReverse(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
