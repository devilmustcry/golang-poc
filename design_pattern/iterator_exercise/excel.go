package main

//TODO: Create an iterator for excel
// 1. when iterate will get next data (Get all data from column first then get next row)
// 2. If no next row and column is found. HasNext function should return false
// 3. main file will only call iterator function

type Excel struct {
	data [][]string
}
