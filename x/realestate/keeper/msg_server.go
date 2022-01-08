package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

type msgServer struct {
	Keeper
}

func (k msgServer) lockAssets(ctx sdk.Context, owner string, fees uint64, bkeeper bankkeeper.Keeper) error {

	addr, _ := sdk.AccAddressFromBech32(owner)
	amount := sdk.NewCoins(sdk.NewCoin(types.GovernanceToken, sdk.NewInt(int64(fees))))

	return bkeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, amount)

}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
