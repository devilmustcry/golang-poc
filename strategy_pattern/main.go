package main

import (
	"fmt"
	"golang-poc/strategy_pattern/factory"
)

func main() {
	input := []int{10, 34, 52, 18429, 1234}
	input2 := 20
	input3 := 1
	f := factory.NewWithdrawStrategyFactory()
	withdrawnStrategy := f.GetWithdrawStrategy(input3)

	fmt.Println(withdrawnStrategy.Withdrawn(input, input2))
}
