package types

const (
	// ModuleName defines the module name
	ModuleName = "el"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_el"
)

var (
	ParamsKey = []byte("p_el")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
