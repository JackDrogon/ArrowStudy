package utils

import "strings"

func stringReverse(src string) string {
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
