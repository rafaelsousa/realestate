package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCertificate = "create_certificate"
	TypeMsgUpdateCertificate = "update_certificate"
	TypeMsgDeleteCertificate = "delete_certificate"
)

var _ sdk.Msg = &MsgCreateCertificate{}

func NewMsgCreateCertificate(creator string, property string, surveyor string, certifiationDate string, certificateText string, propertyMap string) *MsgCreateCertificate {
  return &MsgCreateCertificate{
		Creator: creator,
    Property: property,
    Surveyor: surveyor,
    CertifiationDate: certifiationDate,
    CertificateText: certificateText,
    PropertyMap: propertyMap,
	}
}

func (msg *MsgCreateCertificate) Route() string {
  return RouterKey
}

func (msg *MsgCreateCertificate) Type() string {
  return TypeMsgCreateCertificate
}

func (msg *MsgCreateCertificate) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCertificate) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCertificate) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateCertificate{}

func NewMsgUpdateCertificate(creator string, id uint64, property string, surveyor string, certifiationDate string, certificateText string, propertyMap string) *MsgUpdateCertificate {
  return &MsgUpdateCertificate{
        Id: id,
		Creator: creator,
    Property: property,
    Surveyor: surveyor,
    CertifiationDate: certifiationDate,
    CertificateText: certificateText,
    PropertyMap: propertyMap,
	}
}

func (msg *MsgUpdateCertificate) Route() string {
  return RouterKey
}

func (msg *MsgUpdateCertificate) Type() string {
  return TypeMsgUpdateCertificate
}

func (msg *MsgUpdateCertificate) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCertificate) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCertificate) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteCertificate{}

func NewMsgDeleteCertificate(creator string, id uint64) *MsgDeleteCertificate {
  return &MsgDeleteCertificate{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteCertificate) Route() string {
  return RouterKey
}

func (msg *MsgDeleteCertificate) Type() string {
  return TypeMsgDeleteCertificate
}

func (msg *MsgDeleteCertificate) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCertificate) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCertificate) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
