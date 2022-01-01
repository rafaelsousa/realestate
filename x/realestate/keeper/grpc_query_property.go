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

func (k Keeper) PropertyAll(c context.Context, req *types.QueryAllPropertyRequest) (*types.QueryAllPropertyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var propertys []types.Property
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	propertyStore := prefix.NewStore(store, types.KeyPrefix(types.PropertyKey))

	pageRes, err := query.Paginate(propertyStore, req.Pagination, func(key []byte, value []byte) error {
		var property types.Property
		if err := k.cdc.Unmarshal(value, &property); err != nil {
			return err
		}

		propertys = append(propertys, property)
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

	ctx := sdk.UnwrapSDKContext(c)
	property, found := k.GetProperty(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPropertyResponse{Property: property}, nil
}
