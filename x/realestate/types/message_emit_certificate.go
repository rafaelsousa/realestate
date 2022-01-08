package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEmitCertificate = "emit_certificate"

var _ sdk.Msg = &MsgEmitCertificate{}

func NewMsgEmitCertificate(creator string) *MsgEmitCertificate {
	return &MsgEmitCertificate{
		Creator: creator,
	}
}

func (msg *MsgEmitCertificate) Route() string {
	return RouterKey
}

func (msg *MsgEmitCertificate) Type() string {
	return TypeMsgEmitCertificate
}

func (msg *MsgEmitCertificate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEmitCertificate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEmitCertificate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
