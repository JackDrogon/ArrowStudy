package utils

import "strings"

// asciiStringReverse reverses the input string and returns the result.
// NOTE: only ASCII characters are supported.
func asciiStringReverse(src string) string {
	if len(src) == 0 {
		return ""
	}

	var buf strings.Builder
	buf.Grow(len(src))
	for i := len(src) - 1; i >= 0; i-- {
		buf.WriteByte(src[i])
	}

	return buf.String()
}
