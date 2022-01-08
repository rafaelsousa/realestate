package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) UnlockProperty(goCtx context.Context, msg *types.MsgUnlockProperty) (*types.MsgUnlockPropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUnlockPropertyResponse{}, nil
}
