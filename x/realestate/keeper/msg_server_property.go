package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) CreateProperty(goCtx context.Context, msg *types.MsgCreateProperty) (*types.MsgCreatePropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var property = types.Property{
		Creator:   msg.Creator,
		Address:   msg.Address,
		City:      msg.City,
		State:     msg.State,
		Country:   msg.Country,
		Zipcode:   msg.Zipcode,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
	}

	id := k.AppendProperty(
		ctx,
		property,
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
		Country:   msg.Country,
		Zipcode:   msg.Zipcode,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
	}

	// Checks that the element exists
	val, found := k.GetProperty(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetProperty(ctx, property)

	return &types.MsgUpdatePropertyResponse{}, nil
}

func (k msgServer) DeleteProperty(goCtx context.Context, msg *types.MsgDeleteProperty) (*types.MsgDeletePropertyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetProperty(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProperty(ctx, msg.Id)

	return &types.MsgDeletePropertyResponse{}, nil
}
