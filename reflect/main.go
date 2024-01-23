package main

import (
	"fmt"
	"reflect"
)

type A struct {
	ABC string `json:"abc" validate:"required" validate-msg:"required=400-xxx-001"`
}

func main() {
	o := A{ABC: "12345677"}

	t := reflect.TypeOf(&o).Elem()
	tABC, _ := t.FieldByName("ABC")
	fmt.Println(t)
	fmt.Println(tABC.Tag.Get("xyz"))
}
