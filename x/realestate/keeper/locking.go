package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// GetLockingCount get the total number of locking
func (k Keeper) GetLockingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LockingCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLockingCount set the total number of locking
func (k Keeper) SetLockingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LockingCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLocking appends a locking in the store with a new id and update the count
func (k Keeper) AppendLocking(
	ctx sdk.Context,
	locking types.Locking,
) uint64 {
	// Create the locking
	count := k.GetLockingCount(ctx)

	// Set the ID of the appended value
	locking.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockingKey))
	appendedValue := k.cdc.MustMarshal(&locking)
	store.Set(GetLockingIDBytes(locking.Id), appendedValue)

	// Update locking count
	k.SetLockingCount(ctx, count+1)

	return count
}

// SetLocking set a specific locking in the store
func (k Keeper) SetLocking(ctx sdk.Context, locking types.Locking) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockingKey))
	b := k.cdc.MustMarshal(&locking)
	store.Set(GetLockingIDBytes(locking.Id), b)
}

// GetLocking returns a locking from its id
func (k Keeper) GetLocking(ctx sdk.Context, id uint64) (val types.Locking, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockingKey))
	b := store.Get(GetLockingIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLocking removes a locking from the store
func (k Keeper) RemoveLocking(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockingKey))
	store.Delete(GetLockingIDBytes(id))
}

// GetAllLocking returns all locking
func (k Keeper) GetAllLocking(ctx sdk.Context) (list []types.Locking) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LockingKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Locking
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLockingIDBytes returns the byte representation of the ID
func GetLockingIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLockingIDFromBytes returns ID in uint64 format from a byte array
func GetLockingIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
