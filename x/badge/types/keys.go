package types

const (
	// ModuleName defines the module name
	ModuleName = "badge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_badge"
)

var (
	ParamsKey = []byte("p_badge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
