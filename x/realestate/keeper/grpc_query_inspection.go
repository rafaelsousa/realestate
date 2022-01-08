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

func (k Keeper) InspectionAll(c context.Context, req *types.QueryAllInspectionRequest) (*types.QueryAllInspectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inspections []types.Inspection
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inspectionStore := prefix.NewStore(store, types.KeyPrefix(types.InspectionKey))

	pageRes, err := query.Paginate(inspectionStore, req.Pagination, func(key []byte, value []byte) error {
		var inspection types.Inspection
		if err := k.cdc.Unmarshal(value, &inspection); err != nil {
			return err
		}

		inspections = append(inspections, inspection)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInspectionResponse{Inspection: inspections, Pagination: pageRes}, nil
}

func (k Keeper) Inspection(c context.Context, req *types.QueryGetInspectionRequest) (*types.QueryGetInspectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	inspection, found := k.GetInspection(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetInspectionResponse{Inspection: inspection}, nil
}
