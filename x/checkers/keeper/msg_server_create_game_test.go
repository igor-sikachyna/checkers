package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

func TestCreateGame(t *testing.T) {
	_, msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
