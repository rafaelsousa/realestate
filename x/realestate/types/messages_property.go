package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProperty = "create_property"
	TypeMsgUpdateProperty = "update_property"
	TypeMsgDeleteProperty = "delete_property"
)

var _ sdk.Msg = &MsgCreateProperty{}

func NewMsgCreateProperty(creator string, address string, city string, state string, country string, zipcode string, latitude string, longitude string, owner string) *MsgCreateProperty {
	return &MsgCreateProperty{
		Creator:   creator,
		Address:   address,
		City:      city,
		State:     state,
		Country:   country,
		Zipcode:   zipcode,
		Latitude:  latitude,
		Longitude: longitude,
		Owner:     owner,
	}
}

func (msg *MsgCreateProperty) Route() string {
	return RouterKey
}

func (msg *MsgCreateProperty) Type() string {
	return TypeMsgCreateProperty
}

func (msg *MsgCreateProperty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProperty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProperty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProperty{}

func NewMsgUpdateProperty(creator string, id uint64, address string, city string, state string, country string, zipcode string, latitude string, longitude string) *MsgUpdateProperty {
	return &MsgUpdateProperty{
		Id:        id,
		Creator:   creator,
		Address:   address,
		City:      city,
		State:     state,
		Country:   country,
		Zipcode:   zipcode,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (msg *MsgUpdateProperty) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProperty) Type() string {
	return TypeMsgUpdateProperty
}

func (msg *MsgUpdateProperty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProperty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProperty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProperty{}

func NewMsgDeleteProperty(creator string, id uint64) *MsgDeleteProperty {
	return &MsgDeleteProperty{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteProperty) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProperty) Type() string {
	return TypeMsgDeleteProperty
}

func (msg *MsgDeleteProperty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProperty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProperty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
