package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/igor-sikachyna/checkers/testutil/keeper"
	"github.com/igor-sikachyna/checkers/testutil/nullify"
	"github.com/igor-sikachyna/checkers/x/leaderboard/keeper"
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
)

func createTestLeaderboard(keeper keeper.Keeper, ctx context.Context) types.Leaderboard {
	item := types.Leaderboard{}
	keeper.SetLeaderboard(ctx, item)
	return item
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	rst, found := keeper.GetLeaderboard(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	createTestLeaderboard(keeper, ctx)
	keeper.RemoveLeaderboard(ctx)
	_, found := keeper.GetLeaderboard(ctx)
	require.False(t, found)
}
