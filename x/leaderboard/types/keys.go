package types

const (
	// ModuleName defines the module name
	ModuleName = "leaderboard"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_leaderboard"
)

var (
	ParamsKey = []byte("p_leaderboard")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
