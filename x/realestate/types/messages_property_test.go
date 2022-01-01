package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rafaelsousa/realestate/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateProperty_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateProperty
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateProperty{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateProperty{
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

func TestMsgUpdateProperty_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateProperty
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateProperty{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateProperty{
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

func TestMsgDeleteProperty_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteProperty
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteProperty{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteProperty{
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
