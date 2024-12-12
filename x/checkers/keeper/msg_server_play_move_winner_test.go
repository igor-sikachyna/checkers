package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/igor-sikachyna/checkers/x/checkers/testutil"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

func TestPlayMoveUpToWinner(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)

	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)

	game, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "b",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   40,
		BeforeIndex: "-1",
		AfterIndex:  "-1",
	}, game)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	// 1 event to create a game and 40 moves in total
	require.Len(t, events, 41)
	event := events[len(events)-1]
	require.Equal(t, event.Type, "move-played")
	require.EqualValues(t, sdk.StringEvent{
		Type: "move-played",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: bob},
			{Key: "game-index", Value: "1"},
			{Key: "captured-x", Value: "2"},
			{Key: "captured-y", Value: "5"},
			{Key: "winner", Value: "b"},
			{Key: "board", Value: "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********"},
		},
	}, event)
}
