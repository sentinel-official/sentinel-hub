package v1

func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: Params{
			BlockInterval: 50,
		},
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState()
}

func (m *GenesisState) Validate() error {
	return nil
}
