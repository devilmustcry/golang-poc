package state

type VendingMachine struct {
	State VendingMachineState
}

func NewVendingMachine() VendingMachine {
	return VendingMachine{State: ProductPendingState{}}
}

func (v *VendingMachine) SetState(state VendingMachineState) {
	v.State = state
}

func (v VendingMachine) GetCurrentState() string {
	return v.State.GetState()
}
