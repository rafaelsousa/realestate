package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) UnlockAssets(goCtx context.Context, msg *types.MsgUnlockAssets) (*types.MsgUnlockAssetsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnlockAssetsResponse{}, nil
}
