package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PropertyList:     []Property{},
		CertificateList:  []Certificate{},
		LockingList:      []Locking{},
		InspectionList:   []Inspection{},
		TransferenceList: []Transference{},
		HouseList:        []House{},
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
	// Check for duplicated ID in locking
	lockingIdMap := make(map[uint64]bool)
	lockingCount := gs.GetLockingCount()
	for _, elem := range gs.LockingList {
		if _, ok := lockingIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for locking")
		}
		if elem.Id >= lockingCount {
			return fmt.Errorf("locking id should be lower or equal than the last id")
		}
		lockingIdMap[elem.Id] = true
	}
	// Check for duplicated ID in inspection
	inspectionIdMap := make(map[uint64]bool)
	inspectionCount := gs.GetInspectionCount()
	for _, elem := range gs.InspectionList {
		if _, ok := inspectionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for inspection")
		}
		if elem.Id >= inspectionCount {
			return fmt.Errorf("inspection id should be lower or equal than the last id")
		}
		inspectionIdMap[elem.Id] = true
	}
	// Check for duplicated ID in transference
	transferenceIdMap := make(map[uint64]bool)
	transferenceCount := gs.GetTransferenceCount()
	for _, elem := range gs.TransferenceList {
		if _, ok := transferenceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for transference")
		}
		if elem.Id >= transferenceCount {
			return fmt.Errorf("transference id should be lower or equal than the last id")
		}
		transferenceIdMap[elem.Id] = true
	}
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
