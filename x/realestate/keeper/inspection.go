package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// GetInspectionCount get the total number of inspection
func (k Keeper) GetInspectionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.InspectionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetInspectionCount set the total number of inspection
func (k Keeper) SetInspectionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.InspectionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendInspection appends a inspection in the store with a new id and update the count
func (k Keeper) AppendInspection(
	ctx sdk.Context,
	inspection types.Inspection,
) uint64 {
	// Create the inspection
	count := k.GetInspectionCount(ctx)

	// Set the ID of the appended value
	inspection.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InspectionKey))
	appendedValue := k.cdc.MustMarshal(&inspection)
	store.Set(GetInspectionIDBytes(inspection.Id), appendedValue)

	// Update inspection count
	k.SetInspectionCount(ctx, count+1)

	return count
}

// SetInspection set a specific inspection in the store
func (k Keeper) SetInspection(ctx sdk.Context, inspection types.Inspection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InspectionKey))
	b := k.cdc.MustMarshal(&inspection)
	store.Set(GetInspectionIDBytes(inspection.Id), b)
}

// GetInspection returns a inspection from its id
func (k Keeper) GetInspection(ctx sdk.Context, id uint64) (val types.Inspection, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InspectionKey))
	b := store.Get(GetInspectionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInspection removes a inspection from the store
func (k Keeper) RemoveInspection(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InspectionKey))
	store.Delete(GetInspectionIDBytes(id))
}

// GetAllInspection returns all inspection
func (k Keeper) GetAllInspection(ctx sdk.Context) (list []types.Inspection) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InspectionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Inspection
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetInspectionIDBytes returns the byte representation of the ID
func GetInspectionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetInspectionIDFromBytes returns ID in uint64 format from a byte array
func GetInspectionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
