package keeper_test

import (
	"testing"

	testkeeper "github.com/rafaelsousa/realestate/testutil/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RealestateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
