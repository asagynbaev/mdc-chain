package types

const (
	// ModuleName defines the module name
	ModuleName = "quantumdata"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_quantumdata"
)

var (
	ParamsKey = []byte("p_quantumdata")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
