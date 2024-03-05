package state

import "fmt"

type PaymentPending struct {
}

func (p PaymentPending) SelectProduct(machine *VendingMachine) {
	fmt.Println("Cannot select product please continue to payment")
}
func (p PaymentPending) Pay(machine *VendingMachine) {
	fmt.Println("Payment received... Delivering product")
	machine.SetState(DeliveringProduct{})
}
func (p PaymentPending) DeliverProduct(machine *VendingMachine) {

}

func (p PaymentPending) GetState() string {
	return "Payment pending"
}
