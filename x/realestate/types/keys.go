package types

const (
	// ModuleName defines the module name
	ModuleName = "realestate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_realestate"

	GovernanceToken = "domus"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PropertyKey      = "Property-value-"
	PropertyCountKey = "Property-count-"
)

const (
	CertificateKey      = "Certificate-value-"
	CertificateCountKey = "Certificate-count-"
)

const (
	LockingKey      = "Locking-value-"
	LockingCountKey = "Locking-count-"
)

const (
	InspectionKey      = "Inspection-value-"
	InspectionCountKey = "Inspection-count-"
)

const (
	TransferenceKey      = "Transference-value-"
	TransferenceCountKey = "Transference-count-"
)

const (
	HouseKey      = "House-value-"
	HouseCountKey = "House-count-"
)
