package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func (k msgServer) EmitCertificate(goCtx context.Context, msg *types.MsgEmitCertificate) (*types.MsgEmitCertificateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEmitCertificateResponse{}, nil
}
