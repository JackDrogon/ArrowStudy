package main

import (
	"bytes"
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/csv"
	"github.com/apache/arrow/go/arrow/memory"
)

func csvData() {
	str := `## a simple set of data: int64;float64;string
0;0;str-0
1;1;str-1
2;2;str-2
3;3;str-3
4;4;str-4
5;5;str-5
6;6;str-6
7;7;str-7
8;8;str-8
9;9;str-9
`
	f := bytes.NewBufferString(str)

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
			{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
			{Name: "str", Type: arrow.BinaryTypes.String},
		},
		nil, // no metadata
	)
	csvReader := csv.NewReader(
		f, schema,
		csv.WithComment('#'), csv.WithComma(';'),
		csv.WithChunk(4),
	)
	defer csvReader.Release()

	recordIdx := 0
	for csvReader.Next() {
		record := csvReader.Record()
		for idx, column := range record.Columns() {
			fmt.Printf("record[%d][%q]: %v\n", recordIdx, record.ColumnName(idx), column)
		}
		recordIdx++
	}
}

func memoryTable() {
	mem := memory.NewGoAllocator()
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil, // no metadata
	)

	recordBuilder := array.NewRecordBuilder(mem, schema)
	defer recordBuilder.Release()

	recordBuilder.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{1, 2, 3, 4, 5, 6},
		nil,
	)
	recordBuilder.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{7, 8, 9, 10},
		[]bool{true, true, false, true},
	)
	recordBuilder.Field(1).(*array.Float64Builder).AppendValues(
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		nil,
	)

	record1 := recordBuilder.NewRecord()
	defer record1.Release()

	recordBuilder.Field(0).(*array.Int32Builder).AppendValues([]int32{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)
	recordBuilder.Field(1).(*array.Float64Builder).AppendValues([]float64{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)

	record2 := recordBuilder.NewRecord()
	defer record2.Release()

	table := array.NewTableFromRecords(schema, []array.Record{record1, record2})
	defer table.Release()

	tableReader := array.NewTableReader(table, 5)
	defer tableReader.Release()

	recordIdx := 0
	for tableReader.Next() {
		record := tableReader.Record()
		for idx, column := range record.Columns() {
			fmt.Printf("record[%d][%q]: %v\n", recordIdx, record.ColumnName(idx), column)
		}
		recordIdx++
	}
}

func main() {
	memoryTable()
	println("====================================")
	csvData()
}
