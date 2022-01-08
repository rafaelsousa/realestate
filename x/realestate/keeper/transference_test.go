package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/rafaelsousa/realestate/testutil/keeper"
	"github.com/rafaelsousa/realestate/testutil/nullify"
	"github.com/rafaelsousa/realestate/x/realestate/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/stretchr/testify/require"
)

func createNTransference(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Transference {
	items := make([]types.Transference, n)
	for i := range items {
		items[i].Id = keeper.AppendTransference(ctx, items[i])
	}
	return items
}

func TestTransferenceGet(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNTransference(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTransference(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTransferenceRemove(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNTransference(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTransference(ctx, item.Id)
		_, found := keeper.GetTransference(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTransferenceGetAll(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNTransference(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTransference(ctx)),
	)
}

func TestTransferenceCount(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNTransference(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTransferenceCount(ctx))
}
