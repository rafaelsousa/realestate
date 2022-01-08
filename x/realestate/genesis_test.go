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

		PropertyList: []types.Property{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PropertyCount: 2,
		CertificateList: []types.Certificate{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	CertificateCount: 2,
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RealestateKeeper(t)
	realestate.InitGenesis(ctx, *k, genesisState)
	got := realestate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PropertyList, got.PropertyList)
	require.Equal(t, genesisState.PropertyCount, got.PropertyCount)
	require.ElementsMatch(t, genesisState.CertificateList, got.CertificateList)
require.Equal(t, genesisState.CertificateCount, got.CertificateCount)
// this line is used by starport scaffolding # genesis/test/assert
}
