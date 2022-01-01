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
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PropertyList = k.GetAllProperty(ctx)
	genesis.PropertyCount = k.GetPropertyCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
