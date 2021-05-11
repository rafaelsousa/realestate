package property

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/property/keeper"
	"github.com/rafaelsousa/realestate/x/property/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the owner
	for _, elem := range genState.OwnerList {
		k.SetOwner(ctx, *elem)
	}

	// Set owner count
	k.SetOwnerCount(ctx, uint64(len(genState.OwnerList)))

	// Set all the property
	for _, elem := range genState.PropertyList {
		k.SetProperty(ctx, *elem)
	}

	// Set property count
	k.SetPropertyCount(ctx, uint64(len(genState.PropertyList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all owner
	ownerList := k.GetAllOwner(ctx)
	for _, elem := range ownerList {
		elem := elem
		genesis.OwnerList = append(genesis.OwnerList, &elem)
	}

	// Get all property
	propertyList := k.GetAllProperty(ctx)
	for _, elem := range propertyList {
		elem := elem
		genesis.PropertyList = append(genesis.PropertyList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
