package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/rafaelsousa/realestate/x/realestate/types"
)

func TestPropertyMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateProperty(ctx, &types.MsgCreateProperty{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestPropertyMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateProperty
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateProperty{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateProperty{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateProperty{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateProperty(ctx, &types.MsgCreateProperty{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateProperty(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPropertyMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteProperty
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteProperty{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteProperty{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteProperty{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateProperty(ctx, &types.MsgCreateProperty{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteProperty(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
