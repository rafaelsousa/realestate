package blueprints_test

import (
	"testing"

	keepertest "github.com/rafaelsousa/realestate/testutil/keeper"
	"github.com/rafaelsousa/realestate/testutil/nullify"
	"github.com/rafaelsousa/realestate/x/blueprints"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		HouseList: []types.House{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		HouseCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlueprintsKeeper(t)
	blueprints.InitGenesis(ctx, *k, genesisState)
	got := blueprints.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HouseList, got.HouseList)
	require.Equal(t, genesisState.HouseCount, got.HouseCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
