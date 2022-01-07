package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HouseAll(c context.Context, req *types.QueryAllHouseRequest) (*types.QueryAllHouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var houses []types.House
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	houseStore := prefix.NewStore(store, types.KeyPrefix(types.HouseKey))

	pageRes, err := query.Paginate(houseStore, req.Pagination, func(key []byte, value []byte) error {
		var house types.House
		if err := k.cdc.Unmarshal(value, &house); err != nil {
			return err
		}

		houses = append(houses, house)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHouseResponse{House: houses, Pagination: pageRes}, nil
}

func (k Keeper) House(c context.Context, req *types.QueryGetHouseRequest) (*types.QueryGetHouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	house, found := k.GetHouse(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHouseResponse{House: house}, nil
}
