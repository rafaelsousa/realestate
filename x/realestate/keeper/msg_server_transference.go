package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) CreateTransference(goCtx context.Context, msg *types.MsgCreateTransference) (*types.MsgCreateTransferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var transference = types.Transference{
		Creator:  msg.Creator,
		From:     msg.From,
		To:       msg.To,
		Date:     msg.Date,
		Value:    msg.Value,
		Property: msg.Property,
	}

	id := k.AppendTransference(
		ctx,
		transference,
	)

	return &types.MsgCreateTransferenceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTransference(goCtx context.Context, msg *types.MsgUpdateTransference) (*types.MsgUpdateTransferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var transference = types.Transference{
		Creator:  msg.Creator,
		Id:       msg.Id,
		From:     msg.From,
		To:       msg.To,
		Date:     msg.Date,
		Value:    msg.Value,
		Property: msg.Property,
	}

	// Checks that the element exists
	val, found := k.GetTransference(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTransference(ctx, transference)

	return &types.MsgUpdateTransferenceResponse{}, nil
}

func (k msgServer) DeleteTransference(goCtx context.Context, msg *types.MsgDeleteTransference) (*types.MsgDeleteTransferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTransference(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTransference(ctx, msg.Id)

	return &types.MsgDeleteTransferenceResponse{}, nil
}
