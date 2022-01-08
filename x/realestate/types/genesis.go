package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PropertyList: []Property{},
		CertificateList: []Certificate{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in property
	propertyIdMap := make(map[uint64]bool)
	propertyCount := gs.GetPropertyCount()
	for _, elem := range gs.PropertyList {
		if _, ok := propertyIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for property")
		}
		if elem.Id >= propertyCount {
			return fmt.Errorf("property id should be lower or equal than the last id")
		}
		propertyIdMap[elem.Id] = true
	}
	// Check for duplicated ID in certificate
certificateIdMap := make(map[uint64]bool)
certificateCount := gs.GetCertificateCount()
for _, elem := range gs.CertificateList {
	if _, ok := certificateIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for certificate")
	}
	if elem.Id >= certificateCount {
		return fmt.Errorf("certificate id should be lower or equal than the last id")
	}
	certificateIdMap[elem.Id] = true
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
