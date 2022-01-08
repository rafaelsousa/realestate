package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// GetTransferenceCount get the total number of transference
func (k Keeper) GetTransferenceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TransferenceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTransferenceCount set the total number of transference
func (k Keeper) SetTransferenceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TransferenceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTransference appends a transference in the store with a new id and update the count
func (k Keeper) AppendTransference(
	ctx sdk.Context,
	transference types.Transference,
) uint64 {
	// Create the transference
	count := k.GetTransferenceCount(ctx)

	// Set the ID of the appended value
	transference.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferenceKey))
	appendedValue := k.cdc.MustMarshal(&transference)
	store.Set(GetTransferenceIDBytes(transference.Id), appendedValue)

	// Update transference count
	k.SetTransferenceCount(ctx, count+1)

	return count
}

// SetTransference set a specific transference in the store
func (k Keeper) SetTransference(ctx sdk.Context, transference types.Transference) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferenceKey))
	b := k.cdc.MustMarshal(&transference)
	store.Set(GetTransferenceIDBytes(transference.Id), b)
}

// GetTransference returns a transference from its id
func (k Keeper) GetTransference(ctx sdk.Context, id uint64) (val types.Transference, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferenceKey))
	b := store.Get(GetTransferenceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTransference removes a transference from the store
func (k Keeper) RemoveTransference(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferenceKey))
	store.Delete(GetTransferenceIDBytes(id))
}

// GetAllTransference returns all transference
func (k Keeper) GetAllTransference(ctx sdk.Context) (list []types.Transference) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransferenceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Transference
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTransferenceIDBytes returns the byte representation of the ID
func GetTransferenceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTransferenceIDFromBytes returns ID in uint64 format from a byte array
func GetTransferenceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
