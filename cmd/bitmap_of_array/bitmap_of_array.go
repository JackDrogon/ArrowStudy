package main

import (
	"github.com/JackDrogon/arrow-study/pkg/utils"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	// Step 1: 创建Builder
	int64Builder := array.NewInt64Builder(memory.DefaultAllocator)
	defer int64Builder.Release()

	// Step 2: 通过Builder添加数据
	int64Builder.AppendValues([]int64{1, 2}, nil)
	int64Builder.AppendNull()
	int64Builder.AppendValues([]int64{4, 5, 6, 7, 8, 9, 10}, nil)

	// Step 3: 通过Builder创建Array
	int64Arr := int64Builder.NewArray()
	defer int64Arr.Release()

	// Step 4(optional): 通过Array获取Bitmap
	// 如果一个array没有null元素，那也可以省略bitmap
	// bitmaps := int64Arr.NullBitmapBytes()

	// fmt.Printf("bitmaps : %s\n", utils.BinaryDump(bitmaps)) // fb 03 00 00
	// for _, buf := range int64Arr.Data().Buffers() {
	// 	fmt.Printf("buffer:\n%v", hex.Dump(buf.Buf()))
	// }
	// fmt.Printf("data    : %s\n", int64Arr) // [1 2 (null) 4 5 6 7 8 9 10]
	print(utils.ArrowDump(int64Arr))
}
