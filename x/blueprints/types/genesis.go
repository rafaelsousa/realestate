package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		HouseList: []House{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in house
	houseIdMap := make(map[uint64]bool)
	houseCount := gs.GetHouseCount()
	for _, elem := range gs.HouseList {
		if _, ok := houseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for house")
		}
		if elem.Id >= houseCount {
			return fmt.Errorf("house id should be lower or equal than the last id")
		}
		houseIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
