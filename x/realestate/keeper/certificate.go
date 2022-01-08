package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetCertificateCount get the total number of certificate
func (k Keeper) GetCertificateCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CertificateCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCertificateCount set the total number of certificate
func (k Keeper) SetCertificateCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CertificateCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCertificate appends a certificate in the store with a new id and update the count
func (k Keeper) AppendCertificate(
    ctx sdk.Context,
    certificate types.Certificate,
) uint64 {
	// Create the certificate
    count := k.GetCertificateCount(ctx)

    // Set the ID of the appended value
    certificate.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
    appendedValue := k.cdc.MustMarshal(&certificate)
    store.Set(GetCertificateIDBytes(certificate.Id), appendedValue)

    // Update certificate count
    k.SetCertificateCount(ctx, count+1)

    return count
}

// SetCertificate set a specific certificate in the store
func (k Keeper) SetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	b := k.cdc.MustMarshal(&certificate)
	store.Set(GetCertificateIDBytes(certificate.Id), b)
}

// GetCertificate returns a certificate from its id
func (k Keeper) GetCertificate(ctx sdk.Context, id uint64) (val types.Certificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	b := store.Get(GetCertificateIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCertificate removes a certificate from the store
func (k Keeper) RemoveCertificate(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	store.Delete(GetCertificateIDBytes(id))
}

// GetAllCertificate returns all certificate
func (k Keeper) GetAllCertificate(ctx sdk.Context) (list []types.Certificate) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Certificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetCertificateIDBytes returns the byte representation of the ID
func GetCertificateIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCertificateIDFromBytes returns ID in uint64 format from a byte array
func GetCertificateIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
