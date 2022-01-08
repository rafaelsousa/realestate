package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) RequestCertification(goCtx context.Context, msg *types.MsgRequestCertification) (*types.MsgRequestCertificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestCertificationResponse{}, nil
}
