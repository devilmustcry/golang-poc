package third_party

type ThirdPartyAdapter struct {
	ThirdParty ThirdParty
}

func (t ThirdPartyAdapter) SomeCodeThatWeCanChanged() {
	t.ThirdParty.SomeCodeThatCannotBeChanged()
}
