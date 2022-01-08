package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func TestTransferenceMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTransference(ctx, &types.MsgCreateTransference{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTransferenceMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTransference
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTransference{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTransference{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTransference{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTransference(ctx, &types.MsgCreateTransference{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTransference(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTransferenceMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTransference
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTransference{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTransference{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTransference{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTransference(ctx, &types.MsgCreateTransference{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTransference(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
