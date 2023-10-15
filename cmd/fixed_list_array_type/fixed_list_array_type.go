package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	const N = 3
	var (
		data = [][N]int32{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, -9, -8}}
	)

	listBuilder := array.NewFixedSizeListBuilder(memory.DefaultAllocator, N, arrow.PrimitiveTypes.Int32)
	defer listBuilder.Release()

	valueBuilder := listBuilder.ValueBuilder().(*array.Int32Builder)
	valueBuilder.Reserve(len(data))

	for _, v := range data {
		listBuilder.Append(true)
		valueBuilder.AppendValues(v[:], nil)
	}

	arr := listBuilder.NewArray().(*array.FixedSizeList)
	defer arr.Release()
	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps))

	valueArr := arr.ListValues().(*array.Int32)
	buffers := valueArr.Data().Buffers()

	for _, buf := range buffers {
		fmt.Println(hex.Dump(buf.Buf()))
	}
	fmt.Println(arr)
}
