package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// GetPropertyCount get the total number of property
func (k Keeper) GetPropertyCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PropertyCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPropertyCount set the total number of property
func (k Keeper) SetPropertyCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PropertyCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendProperty appends a property in the store with a new id and update the count
func (k Keeper) AppendProperty(
	ctx sdk.Context,
	property types.Property,
) uint64 {
	// Create the property
	count := k.GetPropertyCount(ctx)

	// Set the ID of the appended value
	property.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	appendedValue := k.cdc.MustMarshal(&property)
	store.Set(GetPropertyIDBytes(property.Id), appendedValue)

	// Update property count
	k.SetPropertyCount(ctx, count+1)

	return count
}

// SetProperty set a specific property in the store
func (k Keeper) SetProperty(ctx sdk.Context, property types.Property) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	b := k.cdc.MustMarshal(&property)
	store.Set(GetPropertyIDBytes(property.Id), b)
}

// GetProperty returns a property from its id
func (k Keeper) GetProperty(ctx sdk.Context, id uint64) (val types.Property, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	b := store.Get(GetPropertyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProperty removes a property from the store
func (k Keeper) RemoveProperty(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	store.Delete(GetPropertyIDBytes(id))
}

// GetAllProperty returns all property
func (k Keeper) GetAllProperty(ctx sdk.Context) (list []types.Property) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Property
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPropertyIDBytes returns the byte representation of the ID
func GetPropertyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPropertyIDFromBytes returns ID in uint64 format from a byte array
func GetPropertyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
