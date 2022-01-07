package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateHouse = "create_house"
	TypeMsgUpdateHouse = "update_house"
	TypeMsgDeleteHouse = "delete_house"
)

var _ sdk.Msg = &MsgCreateHouse{}

func NewMsgCreateHouse(creator string, description string, image string) *MsgCreateHouse {
	return &MsgCreateHouse{
		Creator:     creator,
		Description: description,
		Image:       image,
	}
}

func (msg *MsgCreateHouse) Route() string {
	return RouterKey
}

func (msg *MsgCreateHouse) Type() string {
	return TypeMsgCreateHouse
}

func (msg *MsgCreateHouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHouse{}

func NewMsgUpdateHouse(creator string, id uint64, description string, image string) *MsgUpdateHouse {
	return &MsgUpdateHouse{
		Id:          id,
		Creator:     creator,
		Description: description,
		Image:       image,
	}
}

func (msg *MsgUpdateHouse) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHouse) Type() string {
	return TypeMsgUpdateHouse
}

func (msg *MsgUpdateHouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHouse{}

func NewMsgDeleteHouse(creator string, id uint64) *MsgDeleteHouse {
	return &MsgDeleteHouse{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHouse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHouse) Type() string {
	return TypeMsgDeleteHouse
}

func (msg *MsgDeleteHouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
