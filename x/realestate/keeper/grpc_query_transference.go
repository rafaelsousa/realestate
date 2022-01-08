package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TransferenceAll(c context.Context, req *types.QueryAllTransferenceRequest) (*types.QueryAllTransferenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transferences []types.Transference
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	transferenceStore := prefix.NewStore(store, types.KeyPrefix(types.TransferenceKey))

	pageRes, err := query.Paginate(transferenceStore, req.Pagination, func(key []byte, value []byte) error {
		var transference types.Transference
		if err := k.cdc.Unmarshal(value, &transference); err != nil {
			return err
		}

		transferences = append(transferences, transference)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTransferenceResponse{Transference: transferences, Pagination: pageRes}, nil
}

func (k Keeper) Transference(c context.Context, req *types.QueryGetTransferenceRequest) (*types.QueryGetTransferenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transference, found := k.GetTransference(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTransferenceResponse{Transference: transference}, nil
}
