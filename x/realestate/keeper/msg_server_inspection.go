package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) CreateInspection(goCtx context.Context, msg *types.MsgCreateInspection) (*types.MsgCreateInspectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var inspection = types.Inspection{
		Creator:           msg.Creator,
		Inspector:         msg.Inspector,
		Property:          msg.Property,
		Fees:              msg.Fees,
		InspectionResults: msg.InspectionResults,
	}

	id := k.AppendInspection(
		ctx,
		inspection,
	)

	return &types.MsgCreateInspectionResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateInspection(goCtx context.Context, msg *types.MsgUpdateInspection) (*types.MsgUpdateInspectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var inspection = types.Inspection{
		Creator:           msg.Creator,
		Id:                msg.Id,
		Inspector:         msg.Inspector,
		Property:          msg.Property,
		Fees:              msg.Fees,
		InspectionResults: msg.InspectionResults,
	}

	// Checks that the element exists
	val, found := k.GetInspection(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetInspection(ctx, inspection)

	return &types.MsgUpdateInspectionResponse{}, nil
}

func (k msgServer) DeleteInspection(goCtx context.Context, msg *types.MsgDeleteInspection) (*types.MsgDeleteInspectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetInspection(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveInspection(ctx, msg.Id)

	return &types.MsgDeleteInspectionResponse{}, nil
}
