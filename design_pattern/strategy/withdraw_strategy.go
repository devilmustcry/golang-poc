package strategy

type WithdrawnStrategy interface {
	Withdrawn(wallet []int, amount int) (withdrawn []int)
}
