package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	keepertest "github.com/igor-sikachyna/checkers/testutil/keeper"
	"github.com/igor-sikachyna/checkers/x/checkers/keeper"
	checkers "github.com/igor-sikachyna/checkers/x/checkers/module"
	"github.com/igor-sikachyna/checkers/x/checkers/testutil"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	server, k, context, _, escrow := setupMsgServerCreateGameWithMock(t)
	escrow.ExpectAny(context)
	return server, k, context
}

func setupMsgServerCreateGameWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(k)
	return server, k, ctx, ctrl, bankMock
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1",
	}, *createResponse)
}

func TestCreate1GameHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	game1, found1 := keeper.GetStoredGame(ctx, "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   0,
		BeforeIndex: "-1",
		AfterIndex:  "-1",
		Wager:       45,
	}, game1)
}

func TestCreate1GameGetAll(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	games := keeper.GetAllStoredGame(sdk.UnwrapSDKContext(context))
	require.Len(t, games, 1)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "-1",
		AfterIndex:  "-1",
		Wager:       45,
	}, games[0])
}

func TestCreateGameRedAddressBad(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     "notanaddress",
		Wager:   45,
	})
	require.Nil(t, createResponse)
	require.Equal(t,
		"red address is invalid: notanaddress: decoding bech32 failed: invalid separator index -1",
		err.Error())
}

func TestCreateGameEmptyRedAddress(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     "",
		Wager:   45,
	})
	require.Nil(t, createResponse)
	require.Equal(t,
		"red address is invalid: : empty address string is not allowed",
		err.Error())
}

func TestCreate3Games(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	createResponse2, err2 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   45,
	})
	require.Nil(t, err2)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "2",
	}, *createResponse2)
	createResponse3, err3 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
		Wager:   45,
	})
	require.Nil(t, err3)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "3",
	}, *createResponse3)
}

func TestCreate3GamesHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   45,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
		Wager:   45,
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo)
	game1, found1 := keeper.GetStoredGame(ctx, "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "-1",
		AfterIndex:  "2",
		Wager:       45,
	}, game1)
	game2, found2 := keeper.GetStoredGame(ctx, "2")
	require.True(t, found2)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "1",
		AfterIndex:  "3",
		Wager:       45,
	}, game2)
	game3, found3 := keeper.GetStoredGame(ctx, "3")
	require.True(t, found3)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       alice,
		Red:         bob,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "2",
		AfterIndex:  "-1",
		Wager:       45,
	}, game3)
}

func TestCreate3GamesGetAll(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   45,
	})
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
		Wager:   45,
	})
	games := keeper.GetAllStoredGame(sdk.UnwrapSDKContext(context))
	require.Len(t, games, 3)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "-1",
		AfterIndex:  "2",
		Wager:       45,
	}, games[0])
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "1",
		AfterIndex:  "3",
		Wager:       45,
	}, games[1])
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       alice,
		Red:         bob,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "2",
		AfterIndex:  "-1",
		Wager:       45,
	}, games[2])
}

func TestCreateGameFarFuture(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	systemInfo.NextId = 1024
	keeper.SetSystemInfo(ctx, systemInfo)
	createResponse, err := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "1024",
	}, *createResponse)
	systemInfo, found = keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        1025,
		FifoHeadIndex: "1024",
		FifoTailIndex: "1024",
	}, systemInfo)
	game1, found1 := keeper.GetStoredGame(ctx, "1024")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index:       "1024",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		BeforeIndex: "-1",
		AfterIndex:  "-1",
		Wager:       45,
	}, game1)
}

func TestCreate1GameEmitted(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-game-created",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: alice},
			{Key: "game-index", Value: "1"},
			{Key: "black", Value: bob},
			{Key: "red", Value: carol},
			{Key: "wager", Value: "45"},
		},
	}, event)
}

func TestSavedCreatedDeadlineIsParseable(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	game, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	_, err := game.GetDeadlineAsTime()
	require.Nil(t, err)
}

func TestCreate1GameConsumedGas(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	before := ctx.GasMeter().GasConsumed()
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	after := ctx.GasMeter().GasConsumed()
	require.GreaterOrEqual(t, after, before+25_000)
}
