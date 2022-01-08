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

func createNLocking(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Locking {
	items := make([]types.Locking, n)
	for i := range items {
		items[i].Id = keeper.AppendLocking(ctx, items[i])
	}
	return items
}

func TestLockingGet(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNLocking(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLocking(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLockingRemove(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNLocking(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLocking(ctx, item.Id)
		_, found := keeper.GetLocking(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLockingGetAll(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNLocking(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLocking(ctx)),
	)
}

func TestLockingCount(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNLocking(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLockingCount(ctx))
}
