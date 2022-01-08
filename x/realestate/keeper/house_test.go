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

func createNHouse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.House {
	items := make([]types.House, n)
	for i := range items {
		items[i].Id = keeper.AppendHouse(ctx, items[i])
	}
	return items
}

func TestHouseGet(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNHouse(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHouse(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHouseRemove(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNHouse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHouse(ctx, item.Id)
		_, found := keeper.GetHouse(ctx, item.Id)
		require.False(t, found)
	}
}

func TestHouseGetAll(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNHouse(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHouse(ctx)),
	)
}

func TestHouseCount(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNHouse(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHouseCount(ctx))
}
