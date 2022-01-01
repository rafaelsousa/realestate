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

func createNProperty(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Property {
	items := make([]types.Property, n)
	for i := range items {
		items[i].Id = keeper.AppendProperty(ctx, items[i])
	}
	return items
}

func TestPropertyGet(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNProperty(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProperty(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPropertyRemove(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNProperty(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProperty(ctx, item.Id)
		_, found := keeper.GetProperty(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPropertyGetAll(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNProperty(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProperty(ctx)),
	)
}

func TestPropertyCount(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNProperty(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPropertyCount(ctx))
}
