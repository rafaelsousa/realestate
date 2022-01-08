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

func createNInspection(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Inspection {
	items := make([]types.Inspection, n)
	for i := range items {
		items[i].Id = keeper.AppendInspection(ctx, items[i])
	}
	return items
}

func TestInspectionGet(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNInspection(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetInspection(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestInspectionRemove(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNInspection(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInspection(ctx, item.Id)
		_, found := keeper.GetInspection(ctx, item.Id)
		require.False(t, found)
	}
}

func TestInspectionGetAll(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNInspection(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInspection(ctx)),
	)
}

func TestInspectionCount(t *testing.T) {
	keeper, ctx := keepertest.RealestateKeeper(t)
	items := createNInspection(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetInspectionCount(ctx))
}
