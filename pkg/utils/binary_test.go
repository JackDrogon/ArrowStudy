package utils

import "testing"

func TestBinaryDump(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "single byte",
			input:    []byte{0x01},
			expected: "00000001",
		},
		{
			name:     "multiple bytes",
			input:    []byte{0x01, 0x02, 0x03},
			expected: "00000001 00000010 00000011",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinaryDump(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
