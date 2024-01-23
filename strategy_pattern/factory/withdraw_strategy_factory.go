package factory

import "golang-poc/strategy_pattern/strategy"

type WithdrawStrategyFactory struct {
	mostFirst  strategy.WithdrawnStrategy
	leastFirst strategy.WithdrawnStrategy
}

func NewWithdrawStrategyFactory() *WithdrawStrategyFactory {
	return &WithdrawStrategyFactory{
		mostFirst:  &strategy.MostFirstWithdrawStrategy{},
		leastFirst: &strategy.LeastFirstWithdrawStrategy{},
	}
}

func (w WithdrawStrategyFactory) GetWithdrawStrategy(mode int) strategy.WithdrawnStrategy {
	if mode == 2 {
		return w.mostFirst
	} else {
		return w.leastFirst
	}
}
