package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestCertification = "request_certification"

var _ sdk.Msg = &MsgRequestCertification{}

func NewMsgRequestCertification(creator string) *MsgRequestCertification {
	return &MsgRequestCertification{
		Creator: creator,
	}
}

func (msg *MsgRequestCertification) Route() string {
	return RouterKey
}

func (msg *MsgRequestCertification) Type() string {
	return TypeMsgRequestCertification
}

func (msg *MsgRequestCertification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestCertification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestCertification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
