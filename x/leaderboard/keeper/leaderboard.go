package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
)

// SetLeaderboard set leaderboard in the store
func (k Keeper) SetLeaderboard(ctx context.Context, leaderboard types.Leaderboard) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))
	b := k.cdc.MustMarshal(&leaderboard)
	store.Set([]byte{0}, b)
}

// GetLeaderboard returns leaderboard
func (k Keeper) GetLeaderboard(ctx context.Context) (val types.Leaderboard, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLeaderboard removes leaderboard from the store
func (k Keeper) RemoveLeaderboard(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LeaderboardKey))
	store.Delete([]byte{0})
}
