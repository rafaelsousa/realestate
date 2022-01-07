package blueprints

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/blueprints/keeper"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
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

	genesis.HouseList = k.GetAllHouse(ctx)
	genesis.HouseCount = k.GetHouseCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
