package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/property/types"
)

func (k msgServer) CreateProperty(goCtx context.Context, msg *types.MsgCreateProperty) (*types.MsgCreatePropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendProperty(
		ctx,
		msg.Creator,
		msg.Address,
		msg.City,
		msg.State,
		msg.Zip,
		msg.Country,
		msg.Latitude,
		msg.Longitude,
		msg.Owner,
	)

	return &types.MsgCreatePropertyResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProperty(goCtx context.Context, msg *types.MsgUpdateProperty) (*types.MsgUpdatePropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var property = types.Property{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Address:   msg.Address,
		City:      msg.City,
		State:     msg.State,
		Zip:       msg.Zip,
		Country:   msg.Country,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
		Owner:     msg.Owner,
	}

	// Checks that the element exists
	if !k.HasProperty(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPropertyOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetProperty(ctx, property)

	return &types.MsgUpdatePropertyResponse{}, nil
}

func (k msgServer) DeleteProperty(goCtx context.Context, msg *types.MsgDeleteProperty) (*types.MsgDeletePropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasProperty(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPropertyOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProperty(ctx, msg.Id)

	return &types.MsgDeletePropertyResponse{}, nil
}
