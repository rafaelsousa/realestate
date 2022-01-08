package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func TestInspectionMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateInspection(ctx, &types.MsgCreateInspection{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestInspectionMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateInspection
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateInspection{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateInspection{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateInspection{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateInspection(ctx, &types.MsgCreateInspection{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateInspection(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInspectionMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteInspection
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteInspection{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteInspection{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteInspection{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateInspection(ctx, &types.MsgCreateInspection{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteInspection(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
