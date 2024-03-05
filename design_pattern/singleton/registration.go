package singleton

import (
	"golang-poc/design_pattern/singleton/aimet"
	"sync"
)

type Registration struct {
	Mutex *sync.Mutex
}

func (r Registration) Add() {
	aimet.GetInstance().EmployeeCount++
}
