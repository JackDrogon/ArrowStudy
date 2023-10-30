package main

import (
	"bytes"
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/csv"
)

func readCSV() {
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

func main() {
	readCSV()
}
