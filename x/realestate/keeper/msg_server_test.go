package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/rafaelsousa/realestate/testutil/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RealestateKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
