package third_party

import "fmt"

type ThirdParty struct {
}

func (ThirdParty) SomeCodeThatCannotBeChanged() {
	fmt.Println("We cannot change this code...")
}
