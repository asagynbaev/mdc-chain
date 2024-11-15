package types

const (
	// ModuleName defines the module name
	ModuleName = "mdcchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mdcchain"
)

var (
	ParamsKey = []byte("p_mdcchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
