package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateInspection = "create_inspection"
	TypeMsgUpdateInspection = "update_inspection"
	TypeMsgDeleteInspection = "delete_inspection"
)

var _ sdk.Msg = &MsgCreateInspection{}

func NewMsgCreateInspection(creator string, inspector string, property uint64, fees string, inspectionResults string) *MsgCreateInspection {
	return &MsgCreateInspection{
		Creator:           creator,
		Inspector:         inspector,
		Property:          property,
		Fees:              fees,
		InspectionResults: inspectionResults,
	}
}

func (msg *MsgCreateInspection) Route() string {
	return RouterKey
}

func (msg *MsgCreateInspection) Type() string {
	return TypeMsgCreateInspection
}

func (msg *MsgCreateInspection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInspection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInspection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateInspection{}

func NewMsgUpdateInspection(creator string, id uint64, inspector string, property uint64, fees string, inspectionResults string) *MsgUpdateInspection {
	return &MsgUpdateInspection{
		Id:                id,
		Creator:           creator,
		Inspector:         inspector,
		Property:          property,
		Fees:              fees,
		InspectionResults: inspectionResults,
	}
}

func (msg *MsgUpdateInspection) Route() string {
	return RouterKey
}

func (msg *MsgUpdateInspection) Type() string {
	return TypeMsgUpdateInspection
}

func (msg *MsgUpdateInspection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateInspection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateInspection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteInspection{}

func NewMsgDeleteInspection(creator string, id uint64) *MsgDeleteInspection {
	return &MsgDeleteInspection{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteInspection) Route() string {
	return RouterKey
}

func (msg *MsgDeleteInspection) Type() string {
	return TypeMsgDeleteInspection
}

func (msg *MsgDeleteInspection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteInspection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteInspection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
