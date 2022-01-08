package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnlockProperty = "unlock_property"

var _ sdk.Msg = &MsgUnlockProperty{}

func NewMsgUnlockProperty(creator string) *MsgUnlockProperty {
	return &MsgUnlockProperty{
		Creator: creator,
	}
}

func (msg *MsgUnlockProperty) Route() string {
	return RouterKey
}

func (msg *MsgUnlockProperty) Type() string {
	return TypeMsgUnlockProperty
}

func (msg *MsgUnlockProperty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnlockProperty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnlockProperty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
