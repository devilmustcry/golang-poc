package state

import "fmt"

type DeliveringProduct struct {
}

func (p DeliveringProduct) SelectProduct(machine *VendingMachine) {
	fmt.Println("Cannot select product please continue to payment")
}
func (p DeliveringProduct) Pay(machine *VendingMachine) {
	fmt.Println("Payment received... Delivering product")
}
func (p DeliveringProduct) DeliverProduct(machine *VendingMachine) {
	fmt.Println("Delivered")
	machine.SetState(ProductPendingState{})
}

func (p DeliveringProduct) GetState() string {
	return "Deliver pending"
}
