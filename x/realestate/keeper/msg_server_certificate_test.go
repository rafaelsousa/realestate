package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func TestCertificateMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateCertificate(ctx, &types.MsgCreateCertificate{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestCertificateMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCertificate
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateCertificate{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateCertificate{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateCertificate{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateCertificate(ctx, &types.MsgCreateCertificate{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateCertificate(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCertificateMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCertificate
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteCertificate{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteCertificate{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteCertificate{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateCertificate(ctx, &types.MsgCreateCertificate{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteCertificate(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
