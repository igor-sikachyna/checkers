package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Leaderboard(goCtx context.Context, req *types.QueryGetLeaderboardRequest) (*types.QueryGetLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val := k.GetLeaderboard(ctx)

	return &types.QueryGetLeaderboardResponse{Leaderboard: val}, nil
}
