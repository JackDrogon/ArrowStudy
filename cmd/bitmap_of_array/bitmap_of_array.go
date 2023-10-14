package main

import (
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"

	"github.com/JackDrogon/arrow-study/pkg/utils"
)

func main() {
	// Step 1: 创建Builder
	bldr := array.NewInt64Builder(memory.DefaultAllocator)
	defer bldr.Release()

	// Step 2: 通过Builder添加数据
	bldr.AppendValues([]int64{1, 2}, nil)
	bldr.AppendNull()
	bldr.AppendValues([]int64{4, 5, 6, 7, 8, 9, 10}, nil)

	// Step 3: 通过Builder创建Array
	arr := bldr.NewArray()
	defer arr.Release()

	// Step 4(optional): 通过Array获取Bitmap
	// 如果一个array没有null元素，那也可以省略bitmap
	bitmaps := arr.NullBitmapBytes()

	fmt.Println(utils.BinaryDump(bitmaps)) // fb 03 00 00
	fmt.Println(arr)                       // [1 2 (null) 4 5 6 7 8 9 10]
}
