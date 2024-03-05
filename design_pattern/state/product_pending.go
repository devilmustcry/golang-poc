package state

import (
	"fmt"
)

type ProductPendingState struct {
}

func (p ProductPendingState) SelectProduct(machine *VendingMachine) {
	fmt.Println("Product Selected... Continue to Payment pending state")
	machine.SetState(PaymentPending{})
}
func (p ProductPendingState) Pay(machine *VendingMachine) {
	fmt.Println("Cannot pay please select product first")
}
func (p ProductPendingState) DeliverProduct(machine *VendingMachine) {

	fmt.Println("Cannot deliver please select product first")
}
func (p ProductPendingState) GetState() string {
	return "Product Pending"
}
