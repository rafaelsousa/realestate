package keeper

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	tendermintTime "github.com/tendermint/tendermint/types/time"
)

/*
RequestCertification handles a request for a certification. When first created, a property needs to be inspected and
its limitations defined. The property owner needs, then, request a surveyor to inspect it.

The property states change to "locked" and at this moment, the oMsgLockPropertywner of it is the module account.


*/
func (k msgServer) RequestCertification(goCtx context.Context, msg *types.MsgRequestCertification) (*types.MsgRequestCertificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	// get the property
	property, found := k.GetProperty(ctx, msg.Property)
	if !found {
		resp := types.MsgRequestCertificationResponse{
			Response: "Property not found",
		}
		errorMsg := fmt.Sprintf("The property with the ID %d could not be found in the ledger", msg.Property)
		return &resp, status.Error(codes.InvalidArgument, errorMsg)
	}

	currentTime := tendermintTime.Now()
	timeStr := currentTime.Format(time.RFC850)

	msgLocking := new(types.MsgCreateLocking)

	msgLocking = &types.MsgCreateLocking{
		Creator:     msg.Creator,
		Owner:       property.Owner,
		DateLocking: timeStr,
		Assets:      msg.Fees,
		Property:    property.Id,
	}

	//Get the module account
	err := k.lockAssets(ctx, property.Owner, msg.Fees, k.bankKeeper)
	if err != nil {
		return nil, err
	}

	_, err = k.CreateLocking(goCtx, msgLocking)
	if err != nil {
		errorMsg := fmt.Sprintf("Could not lock property with ID %d ", msg.Property)
		resp := types.MsgRequestCertificationResponse{
			Response: "Property not found",
		}
		return &resp, status.Error(codes.Internal, errorMsg)
	}

	return &types.MsgRequestCertificationResponse{}, nil
}
