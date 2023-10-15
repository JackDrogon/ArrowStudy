package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
)

// TODO(Drogon): think fmt to string.Buffer direct
// ArrowDump returns a string representation of the given arrow.Array.
// If the array is nil, an empty string is returned.
// The function first dumps the null bitmap, then the buffers, and finally the data.
// The null bitmap is dumped as a binary string.
// The buffers are dumped as hexadecimal strings.
// The data is dumped using the default string representation of the array.
func ArrowDump(arr arrow.Array) string {
	if arr == nil {
		return ""
	}

	var stringBuilder strings.Builder
	var s string
	bitmaps := arr.NullBitmapBytes()

	// Step 1: dump bitmap
	s = fmt.Sprintf("bitmaps : %s\n", BinaryDump(bitmaps))
	stringBuilder.WriteString(s)

	// Step 2: dump buffer
	buffers := arr.Data().Buffers()
	stringBuilder.WriteString("Buffers:\n")
	for idx, buf := range buffers {
		stringBuilder.WriteString(fmt.Sprintf("buffer %d:\n", idx))
		stringBuilder.WriteString(hex.Dump(buf.Buf()))
	}

	// Step 3: dump data
	s = fmt.Sprintf("data    : %s\n", arr)
	stringBuilder.WriteString(s)

	return stringBuilder.String()
}
