package singleton

import (
	"golang-poc/design_pattern/singleton/aimet"
	"sync"
)

type EmployeeStatisticCount struct {
	Mutex *sync.Mutex
}

func (a EmployeeStatisticCount) Multiply() {
	aimet.GetInstance().EmployeeCount *= 2
}
