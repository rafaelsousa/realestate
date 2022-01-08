package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnlockAssets = "unlock_assets"

var _ sdk.Msg = &MsgUnlockAssets{}

func NewMsgUnlockAssets(creator string) *MsgUnlockAssets {
	return &MsgUnlockAssets{
		Creator: creator,
	}
}

func (msg *MsgUnlockAssets) Route() string {
	return RouterKey
}

func (msg *MsgUnlockAssets) Type() string {
	return TypeMsgUnlockAssets
}

func (msg *MsgUnlockAssets) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnlockAssets) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnlockAssets) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
