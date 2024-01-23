package strategy

import "sort"

type MostFirstWithdrawStrategy struct {
}

func (m MostFirstWithdrawStrategy) Withdrawn(wallet []int, amount int) (withdrawn []int) {
	sort.Slice(wallet, func(i, j int) bool {
		return wallet[i] > wallet[j]
	})
	acc := 0
	withdrawn = make([]int, 0)
	for _, money := range wallet {
		acc = acc + money
		withdrawn = append(withdrawn, money)
		if acc >= amount {
			break
		}
	}
	return withdrawn
}
