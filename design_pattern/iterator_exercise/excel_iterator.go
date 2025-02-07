package main

type ExcelIterator struct {
	i int
	j int
	e Excel
}

func (ite *ExcelIterator) HasNext() bool {
	if ite.i < len(ite.e.data) {
		return true
	}
	return false
}

func (ite *ExcelIterator) Next() string {
	if ite.HasNext() {
		data := ite.e.data[ite.i][ite.j]
		ite.j++
		if ite.j >= len(ite.e.data[ite.i]) {
			ite.i++
			ite.j = 0
		}
		return data
	}
	return "out"
}
