package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) CreateLocking(goCtx context.Context, msg *types.MsgCreateLocking) (*types.MsgCreateLockingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var locking = types.Locking{
		Creator:       msg.Creator,
		Owner:         msg.Owner,
		DateLocking:   msg.DateLocking,
		DateUnlocking: msg.DateUnlocking,
		Assets:        msg.Assets,
		Property:      msg.Property,
	}

	id := k.AppendLocking(
		ctx,
		locking,
	)

	return &types.MsgCreateLockingResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateLocking(goCtx context.Context, msg *types.MsgUpdateLocking) (*types.MsgUpdateLockingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var locking = types.Locking{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Owner:         msg.Owner,
		DateLocking:   msg.DateLocking,
		DateUnlocking: msg.DateUnlocking,
		Assets:        msg.Assets,
		Property:      msg.Property,
	}

	// Checks that the element exists
	val, found := k.GetLocking(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetLocking(ctx, locking)

	return &types.MsgUpdateLockingResponse{}, nil
}

func (k msgServer) DeleteLocking(goCtx context.Context, msg *types.MsgDeleteLocking) (*types.MsgDeleteLockingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetLocking(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLocking(ctx, msg.Id)

	return &types.MsgDeleteLockingResponse{}, nil
}
