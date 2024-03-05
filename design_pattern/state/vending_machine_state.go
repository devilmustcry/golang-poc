package state

type VendingMachineState interface {
	SelectProduct(machine *VendingMachine)
	Pay(machine *VendingMachine)
	DeliverProduct(machine *VendingMachine)
	GetState() string
}
