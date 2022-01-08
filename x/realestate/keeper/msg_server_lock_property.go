package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) LockProperty(goCtx context.Context, msg *types.MsgLockProperty) (*types.MsgLockPropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLockPropertyResponse{}, nil
}
