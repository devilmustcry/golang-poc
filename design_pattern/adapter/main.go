package main

import (
	"golang-poc/design_pattern/adapter/first_party"
	"golang-poc/design_pattern/adapter/third_party"
)

func main() {
	var party1 first_party.Party
	var party2 first_party.Party
	party1 = first_party.FirstParty{}
	party1.SomeCodeThatWeCanChanged()

	party2 = third_party.ThirdPartyAdapter{
		ThirdParty: third_party.ThirdParty{},
	}

	party2.SomeCodeThatWeCanChanged()

}
