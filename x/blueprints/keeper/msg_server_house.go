package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
)

func (k msgServer) CreateHouse(goCtx context.Context, msg *types.MsgCreateHouse) (*types.MsgCreateHouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var house = types.House{
		Creator:     msg.Creator,
		Description: msg.Description,
		Image:       msg.Image,
	}

	id := k.AppendHouse(
		ctx,
		house,
	)

	return &types.MsgCreateHouseResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHouse(goCtx context.Context, msg *types.MsgUpdateHouse) (*types.MsgUpdateHouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var house = types.House{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Description: msg.Description,
		Image:       msg.Image,
	}

	// Checks that the element exists
	val, found := k.GetHouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHouse(ctx, house)

	return &types.MsgUpdateHouseResponse{}, nil
}

func (k msgServer) DeleteHouse(goCtx context.Context, msg *types.MsgDeleteHouse) (*types.MsgDeleteHouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetHouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHouse(ctx, msg.Id)

	return &types.MsgDeleteHouseResponse{}, nil
}
