package realestate_test

import (
	"testing"

	keepertest "github.com/rafaelsousa/realestate/testutil/keeper"
	"github.com/rafaelsousa/realestate/testutil/nullify"
	"github.com/rafaelsousa/realestate/x/realestate"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RealestateKeeper(t)
	realestate.InitGenesis(ctx, *k, genesisState)
	got := realestate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
