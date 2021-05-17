package property

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/property/keeper"
	"github.com/rafaelsousa/realestate/x/property/types"
)

func handleMsgCreateProperty(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateProperty) (*sdk.Result, error) {
	k.CreateProperty(ctx, *msg)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}