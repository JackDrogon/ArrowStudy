package utils

import (
	"fmt"
	"strings"
)

// BinaryDump returns a string representation of the input byte slice in binary format.
// If the input slice is empty, an empty string is returned.
// The returned string is constructed by concatenating the binary representation of each byte in the input slice.
// Each byte is represented as an 8-bit binary string.
// The function returns the resulting string.
func BinaryDump(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	var buf strings.Builder
	buf.Grow(len(data) * 9) // 8 for binary and 1 for space
	for i, n := range data {
		if i != 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(asciiStringReverse(fmt.Sprintf("%08b", n)))
	}

	return buf.String()
}
