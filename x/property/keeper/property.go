package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/property/types"
	"strconv"
)

// GetPropertyCount get the total number of property
func (k Keeper) GetPropertyCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyCountKey))
	byteKey := types.KeyPrefix(types.PropertyCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetPropertyCount set the total number of property
func (k Keeper) SetPropertyCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyCountKey))
	byteKey := types.KeyPrefix(types.PropertyCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendProperty appends a property in the store with a new id and update the count
func (k Keeper) AppendProperty(
	ctx sdk.Context,
	creator string,
	address string,
	city string,
	state string,
	zip string,
	country string,
	latitude string,
	longitude string,
	owner_addr string,
) uint64 {
	// Create the property
	count := k.GetPropertyCount(ctx)
	var property = types.Property{
		Creator:    creator,
		Id:         count,
		Address:    address,
		City:       city,
		State:      state,
		Zip:        zip,
		Country:    country,
		Latitude:   latitude,
		Longitude:  longitude,
		Owneraddr: owner_addr,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	value := k.cdc.MustMarshalBinaryBare(&property)
	store.Set(GetPropertyIDBytes(property.Id), value)

	// Update property count
	k.SetPropertyCount(ctx, count+1)

	return count
}

// SetProperty set a specific property in the store
func (k Keeper) SetProperty(ctx sdk.Context, property types.Property) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	b := k.cdc.MustMarshalBinaryBare(&property)
	store.Set(GetPropertyIDBytes(property.Id), b)
}

// GetProperty returns a property from its id
func (k Keeper) GetProperty(ctx sdk.Context, id uint64) types.Property {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	var property types.Property
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPropertyIDBytes(id)), &property)
	return property
}

// HasProperty checks if the property exists in the store
func (k Keeper) HasProperty(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PropertyKey))
	return store.Has(GetPropertyIDBytes(id))
}

// GetPropertyOwner returns the creator of the property
func (k Keeper) GetPropertyOwner(ctx sdk.Context, id uint64) string {
	return k.GetProperty(ctx, id).Creator
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
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
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
