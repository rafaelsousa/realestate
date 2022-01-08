package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/rafaelsousa/realestate/testutil/sample"
)

func TestMsgCreateCertificate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCertificate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCertificate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCertificate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateCertificate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCertificate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCertificate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCertificate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteCertificate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCertificate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCertificate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCertificate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
