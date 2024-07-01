package v3

import (
	"fmt"

	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func NewGenesisState(nodes []v2.Node, params Params) *GenesisState {
	return &GenesisState{
		Nodes:  nodes,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	m := make(map[string]bool)
	for _, node := range state.Nodes {
		if m[node.Address] {
			return fmt.Errorf("found a duplicate node for address %s", node.Address)
		}

		m[node.Address] = true

		if err := node.Validate(); err != nil {
			return err
		}
	}

	return nil
}
