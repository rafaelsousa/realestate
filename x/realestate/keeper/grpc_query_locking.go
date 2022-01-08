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

func (k Keeper) LockingAll(c context.Context, req *types.QueryAllLockingRequest) (*types.QueryAllLockingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lockings []types.Locking
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	lockingStore := prefix.NewStore(store, types.KeyPrefix(types.LockingKey))

	pageRes, err := query.Paginate(lockingStore, req.Pagination, func(key []byte, value []byte) error {
		var locking types.Locking
		if err := k.cdc.Unmarshal(value, &locking); err != nil {
			return err
		}

		lockings = append(lockings, locking)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLockingResponse{Locking: lockings, Pagination: pageRes}, nil
}

func (k Keeper) Locking(c context.Context, req *types.QueryGetLockingRequest) (*types.QueryGetLockingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	locking, found := k.GetLocking(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLockingResponse{Locking: locking}, nil
}
