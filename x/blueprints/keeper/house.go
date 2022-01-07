package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
)

// GetHouseCount get the total number of house
func (k Keeper) GetHouseCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HouseCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHouseCount set the total number of house
func (k Keeper) SetHouseCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.HouseCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHouse appends a house in the store with a new id and update the count
func (k Keeper) AppendHouse(
	ctx sdk.Context,
	house types.House,
) uint64 {
	// Create the house
	count := k.GetHouseCount(ctx)

	// Set the ID of the appended value
	house.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HouseKey))
	appendedValue := k.cdc.MustMarshal(&house)
	store.Set(GetHouseIDBytes(house.Id), appendedValue)

	// Update house count
	k.SetHouseCount(ctx, count+1)

	return count
}

// SetHouse set a specific house in the store
func (k Keeper) SetHouse(ctx sdk.Context, house types.House) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HouseKey))
	b := k.cdc.MustMarshal(&house)
	store.Set(GetHouseIDBytes(house.Id), b)
}

// GetHouse returns a house from its id
func (k Keeper) GetHouse(ctx sdk.Context, id uint64) (val types.House, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HouseKey))
	b := store.Get(GetHouseIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHouse removes a house from the store
func (k Keeper) RemoveHouse(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HouseKey))
	store.Delete(GetHouseIDBytes(id))
}

// GetAllHouse returns all house
func (k Keeper) GetAllHouse(ctx sdk.Context) (list []types.House) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HouseKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.House
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHouseIDBytes returns the byte representation of the ID
func GetHouseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHouseIDFromBytes returns ID in uint64 format from a byte array
func GetHouseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
