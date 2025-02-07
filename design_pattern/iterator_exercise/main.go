package main

import "fmt"

func main() {
	//TODO: Call excel iterator here

	data := [][]string{{"a", "b", "c"}, {"1", "2", "3"}}

	excel := Excel{data: data}

	excelIterator := ExcelIterator{
		e: excel,
	}

	for excelIterator.HasNext() {
		fmt.Println(excelIterator.Next())
	}
}
