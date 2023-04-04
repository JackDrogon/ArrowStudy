package main

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

func main() {
	mem := memory.NewGoAllocator()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil, // no metadata
	)

	b := array.NewRecordBuilder(mem, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{1, 2, 3, 4, 5, 6},
		nil,
	)
	b.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{7, 8, 9, 10},
		[]bool{true, true, false, true},
	)
	b.Field(1).(*array.Float64Builder).AppendValues(
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		nil,
	)

	rec1 := b.NewRecord()
	defer rec1.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)

	rec2 := b.NewRecord()
	defer rec2.Release()

	tbl := array.NewTableFromRecords(schema, []array.Record{rec1, rec2})
	defer tbl.Release()

	tr := array.NewTableReader(tbl, 5)
	defer tr.Release()

	n := 0
	for tr.Next() {
		rec := tr.Record()
		for i, col := range rec.Columns() {
			fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
		}
		n++
	}
}
