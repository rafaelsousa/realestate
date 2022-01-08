package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLockProperty = "lock_property"

var _ sdk.Msg = &MsgLockProperty{}

func NewMsgLockProperty(creator string) *MsgLockProperty {
	return &MsgLockProperty{
		Creator: creator,
	}
}

func (msg *MsgLockProperty) Route() string {
	return RouterKey
}

func (msg *MsgLockProperty) Type() string {
	return TypeMsgLockProperty
}

func (msg *MsgLockProperty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLockProperty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLockProperty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
