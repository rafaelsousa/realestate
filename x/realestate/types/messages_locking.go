package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateLocking = "create_locking"
	TypeMsgUpdateLocking = "update_locking"
	TypeMsgDeleteLocking = "delete_locking"
)

var _ sdk.Msg = &MsgCreateLocking{}

func NewMsgCreateLocking(creator string, owner string, dateLocking string, dateUnlocking string, assets *sdk.Coin, property uint64) *MsgCreateLocking {
	return &MsgCreateLocking{
		Creator:       creator,
		Owner:         owner,
		DateLocking:   dateLocking,
		DateUnlocking: dateUnlocking,
		Assets:        assets,
		Property:      property,
	}
}

func (msg *MsgCreateLocking) Route() string {
	return RouterKey
}

func (msg *MsgCreateLocking) Type() string {
	return TypeMsgCreateLocking
}

func (msg *MsgCreateLocking) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLocking) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLocking) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLocking{}

func NewMsgUpdateLocking(creator string, id uint64, owner string, dateLocking string, dateUnlocking string, assets sdk.Coin, property uint64) *MsgUpdateLocking {
	return &MsgUpdateLocking{
		Id:            id,
		Creator:       creator,
		Owner:         owner,
		DateLocking:   dateLocking,
		DateUnlocking: dateUnlocking,
		Assets:        &assets,
		Property:      property,
	}
}

func (msg *MsgUpdateLocking) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLocking) Type() string {
	return TypeMsgUpdateLocking
}

func (msg *MsgUpdateLocking) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLocking) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLocking) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLocking{}

func NewMsgDeleteLocking(creator string, id uint64) *MsgDeleteLocking {
	return &MsgDeleteLocking{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteLocking) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLocking) Type() string {
	return TypeMsgDeleteLocking
}

func (msg *MsgDeleteLocking) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLocking) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLocking) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
