package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTransference = "create_transference"
	TypeMsgUpdateTransference = "update_transference"
	TypeMsgDeleteTransference = "delete_transference"
)

var _ sdk.Msg = &MsgCreateTransference{}

func NewMsgCreateTransference(creator string, from string, to string, date string, value string, property uint64) *MsgCreateTransference {
	return &MsgCreateTransference{
		Creator:  creator,
		From:     from,
		To:       to,
		Date:     date,
		Value:    value,
		Property: property,
	}
}

func (msg *MsgCreateTransference) Route() string {
	return RouterKey
}

func (msg *MsgCreateTransference) Type() string {
	return TypeMsgCreateTransference
}

func (msg *MsgCreateTransference) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTransference) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTransference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTransference{}

func NewMsgUpdateTransference(creator string, id uint64, from string, to string, date string, value string, property uint64) *MsgUpdateTransference {
	return &MsgUpdateTransference{
		Id:       id,
		Creator:  creator,
		From:     from,
		To:       to,
		Date:     date,
		Value:    value,
		Property: property,
	}
}

func (msg *MsgUpdateTransference) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTransference) Type() string {
	return TypeMsgUpdateTransference
}

func (msg *MsgUpdateTransference) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTransference) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTransference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTransference{}

func NewMsgDeleteTransference(creator string, id uint64) *MsgDeleteTransference {
	return &MsgDeleteTransference{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTransference) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTransference) Type() string {
	return TypeMsgDeleteTransference
}

func (msg *MsgDeleteTransference) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTransference) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTransference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
