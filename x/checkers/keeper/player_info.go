package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

// SetPlayerInfo set a specific playerInfo in the store from its index
func (k Keeper) SetPlayerInfo(ctx context.Context, playerInfo types.PlayerInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerInfoKeyPrefix))
	b := k.cdc.MustMarshal(&playerInfo)
	store.Set(types.PlayerInfoKey(
		playerInfo.Index,
	), b)
}

// GetPlayerInfo returns a playerInfo from its index
func (k Keeper) GetPlayerInfo(
	ctx context.Context,
	index string,

) (val types.PlayerInfo, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerInfoKeyPrefix))

	b := store.Get(types.PlayerInfoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlayerInfo removes a playerInfo from the store
func (k Keeper) RemovePlayerInfo(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerInfoKeyPrefix))
	store.Delete(types.PlayerInfoKey(
		index,
	))
}

// GetAllPlayerInfo returns all playerInfo
func (k Keeper) GetAllPlayerInfo(ctx context.Context) (list []types.PlayerInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerInfoKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PlayerInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
