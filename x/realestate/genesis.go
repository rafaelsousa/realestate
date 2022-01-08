package realestate

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the property
	for _, elem := range genState.PropertyList {
		k.SetProperty(ctx, elem)
	}

	// Set property count
	k.SetPropertyCount(ctx, genState.PropertyCount)
	// Set all the certificate
	for _, elem := range genState.CertificateList {
		k.SetCertificate(ctx, elem)
	}

	// Set certificate count
	k.SetCertificateCount(ctx, genState.CertificateCount)
	// Set all the locking
	for _, elem := range genState.LockingList {
		k.SetLocking(ctx, elem)
	}

	// Set locking count
	k.SetLockingCount(ctx, genState.LockingCount)
	// Set all the inspection
	for _, elem := range genState.InspectionList {
		k.SetInspection(ctx, elem)
	}

	// Set inspection count
	k.SetInspectionCount(ctx, genState.InspectionCount)
	// Set all the transference
	for _, elem := range genState.TransferenceList {
		k.SetTransference(ctx, elem)
	}

	// Set transference count
	k.SetTransferenceCount(ctx, genState.TransferenceCount)
	// Set all the house
	for _, elem := range genState.HouseList {
		k.SetHouse(ctx, elem)
	}

	// Set house count
	k.SetHouseCount(ctx, genState.HouseCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PropertyList = k.GetAllProperty(ctx)
	genesis.PropertyCount = k.GetPropertyCount(ctx)
	genesis.CertificateList = k.GetAllCertificate(ctx)
	genesis.CertificateCount = k.GetCertificateCount(ctx)
	genesis.LockingList = k.GetAllLocking(ctx)
	genesis.LockingCount = k.GetLockingCount(ctx)
	genesis.InspectionList = k.GetAllInspection(ctx)
	genesis.InspectionCount = k.GetInspectionCount(ctx)
	genesis.TransferenceList = k.GetAllTransference(ctx)
	genesis.TransferenceCount = k.GetTransferenceCount(ctx)
	genesis.HouseList = k.GetAllHouse(ctx)
	genesis.HouseCount = k.GetHouseCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
