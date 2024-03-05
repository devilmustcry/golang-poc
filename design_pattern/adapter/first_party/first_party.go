package first_party

import "fmt"

type Party interface {
	SomeCodeThatWeCanChanged()
}

type FirstParty struct {
}

func (FirstParty) SomeCodeThatWeCanChanged() {
	fmt.Println("Example code")
}
