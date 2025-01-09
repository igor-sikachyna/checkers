package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PlayerInfoAll(ctx context.Context, req *types.QueryAllPlayerInfoRequest) (*types.QueryAllPlayerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var playerInfos []types.PlayerInfo

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	playerInfoStore := prefix.NewStore(store, types.KeyPrefix(types.PlayerInfoKeyPrefix))

	pageRes, err := query.Paginate(playerInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var playerInfo types.PlayerInfo
		if err := k.cdc.Unmarshal(value, &playerInfo); err != nil {
			return err
		}

		playerInfos = append(playerInfos, playerInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPlayerInfoResponse{PlayerInfo: playerInfos, Pagination: pageRes}, nil
}

func (k Keeper) PlayerInfo(ctx context.Context, req *types.QueryGetPlayerInfoRequest) (*types.QueryGetPlayerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPlayerInfo(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPlayerInfoResponse{PlayerInfo: val}, nil
}
