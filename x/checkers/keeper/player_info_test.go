package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/igor-sikachyna/checkers/testutil/keeper"
	"github.com/igor-sikachyna/checkers/testutil/nullify"
	"github.com/igor-sikachyna/checkers/x/checkers/keeper"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPlayerInfo(keeper keeper.Keeper, ctx context.Context, n int) []types.PlayerInfo {
	items := make([]types.PlayerInfo, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPlayerInfo(ctx, items[i])
	}
	return items
}

func TestPlayerInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlayerInfo(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPlayerInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlayerInfo(ctx,
			item.Index,
		)
		_, found := keeper.GetPlayerInfo(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPlayerInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlayerInfo(ctx)),
	)
}
