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
		LockingList: []types.Locking{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LockingCount: 2,
		InspectionList: []types.Inspection{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		InspectionCount: 2,
		TransferenceList: []types.Transference{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TransferenceCount: 2,
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
	require.ElementsMatch(t, genesisState.LockingList, got.LockingList)
	require.Equal(t, genesisState.LockingCount, got.LockingCount)
	require.ElementsMatch(t, genesisState.InspectionList, got.InspectionList)
	require.Equal(t, genesisState.InspectionCount, got.InspectionCount)
	require.ElementsMatch(t, genesisState.TransferenceList, got.TransferenceList)
	require.Equal(t, genesisState.TransferenceCount, got.TransferenceCount)
	require.ElementsMatch(t, genesisState.HouseList, got.HouseList)
	require.Equal(t, genesisState.HouseCount, got.HouseCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
