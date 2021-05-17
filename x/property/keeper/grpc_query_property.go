package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rafaelsousa/realestate/x/property/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PropertyAll(c context.Context, req *types.QueryAllPropertyRequest) (*types.QueryAllPropertyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var propertys []*types.Property
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	propertyStore := prefix.NewStore(store, types.KeyPrefix(types.PropertyKey))

	pageRes, err := query.Paginate(propertyStore, req.Pagination, func(key []byte, value []byte) error {
		var property types.Property
		if err := k.cdc.UnmarshalBinaryBare(value, &property); err != nil {
			return err
		}

		propertys = append(propertys, &property)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPropertyResponse{Property: propertys, Pagination: pageRes}, nil
}

func (k Keeper) Property(c context.Context, req *types.QueryGetPropertyRequest) (*types.QueryGetPropertyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var property types.Property
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasProperty(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPropertyIDBytes(req.Id)), &property)

	return &types.QueryGetPropertyResponse{Property: &property}, nil
}
