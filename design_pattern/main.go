package main

import (
	"fmt"
	"golang-poc/design_pattern/state"
)

func main() {

	v := state.NewVendingMachine()
	v.State.DeliverProduct(&v)
	fmt.Println(v.GetCurrentState())
	v.State.SelectProduct(&v)
	fmt.Println(v.GetCurrentState())
	v.State.Pay(&v)
	fmt.Println(v.GetCurrentState())
	v.State.DeliverProduct(&v)
	fmt.Println(v.GetCurrentState())

}
