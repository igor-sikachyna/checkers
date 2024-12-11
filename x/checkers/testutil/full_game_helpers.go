package testutil

import (
	"context"
	"testing"

	"github.com/igor-sikachyna/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

type GameMoveTest struct {
	player string
	fromX  uint64
	fromY  uint64
	toX    uint64
	toY    uint64
}

var (
	Game1Moves = []GameMoveTest{
		{"b", 1, 2, 2, 3}, // "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"
		{"r", 0, 5, 1, 4}, // "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*"
		{"b", 2, 3, 0, 5}, // "*b*b*b*b|b*b*b*b*|***b*b*b|********|********|b*r*r*r*|*r*r*r*r|r*r*r*r*"
		{"r", 4, 5, 3, 4}, // "*b*b*b*b|b*b*b*b*|***b*b*b|********|***r****|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"b", 3, 2, 2, 3}, // "*b*b*b*b|b*b*b*b*|*****b*b|**b*****|***r****|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"r", 3, 4, 1, 2}, // "*b*b*b*b|b*b*b*b*|*r***b*b|********|********|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"b", 0, 1, 2, 3}, // "*b*b*b*b|**b*b*b*|*****b*b|**b*****|********|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"r", 2, 5, 3, 4}, // "*b*b*b*b|**b*b*b*|*****b*b|**b*****|***r****|b*****r*|*r*r*r*r|r*r*r*r*"
		{"b", 2, 3, 4, 5}, // "*b*b*b*b|**b*b*b*|*****b*b|********|********|b***b*r*|*r*r*r*r|r*r*r*r*"
		{"r", 5, 6, 3, 4}, // "*b*b*b*b|**b*b*b*|*****b*b|********|***r****|b*****r*|*r*r***r|r*r*r*r*"
		{"b", 5, 2, 4, 3}, // "*b*b*b*b|**b*b*b*|*******b|****b***|***r****|b*****r*|*r*r***r|r*r*r*r*"
		{"r", 3, 4, 5, 2}, // "*b*b*b*b|**b*b*b*|*****r*b|********|********|b*****r*|*r*r***r|r*r*r*r*"
		{"b", 6, 1, 4, 3}, // "*b*b*b*b|**b*b***|*******b|****b***|********|b*****r*|*r*r***r|r*r*r*r*"
		{"r", 6, 5, 5, 4}, // "*b*b*b*b|**b*b***|*******b|****b***|*****r**|b*******|*r*r***r|r*r*r*r*"
		{"b", 4, 3, 6, 5}, // "*b*b*b*b|**b*b***|*******b|********|********|b*****b*|*r*r***r|r*r*r*r*"
		{"r", 7, 6, 5, 4}, // "*b*b*b*b|**b*b***|*******b|********|*****r**|b*******|*r*r****|r*r*r*r*"
		{"b", 7, 2, 6, 3}, // "*b*b*b*b|**b*b***|********|******b*|*****r**|b*******|*r*r****|r*r*r*r*"
		{"r", 5, 4, 7, 2}, // "*b*b*b*b|**b*b***|*******r|********|********|b*******|*r*r****|r*r*r*r*"
		{"b", 4, 1, 3, 2}, // "*b*b*b*b|**b*****|***b***r|********|********|b*******|*r*r****|r*r*r*r*"
		{"r", 3, 6, 4, 5}, // "*b*b*b*b|**b*****|***b***r|********|********|b***r***|*r******|r*r*r*r*"
		{"b", 5, 0, 4, 1}, // "*b*b***b|**b*b***|***b***r|********|********|b***r***|*r******|r*r*r*r*"
		{"r", 2, 7, 3, 6}, // "*b*b***b|**b*b***|***b***r|********|********|b***r***|*r*r****|r***r*r*"
		{"b", 0, 5, 2, 7}, // "*b*b***b|**b*b***|***b***r|********|********|****r***|***r****|r*B*r*r*"
		{"r", 4, 5, 3, 4}, // "*b*b***b|**b*b***|***b***r|********|***r****|********|***r****|r*B*r*r*"
		{"b", 2, 7, 4, 5}, // "*b*b***b|**b*b***|***b***r|********|***r****|****B***|********|r***r*r*"
		// Captures again
		{"b", 4, 5, 2, 3}, // "*b*b***b|**b*b***|***b***r|**B*****|********|********|********|r***r*r*"
		{"r", 6, 7, 5, 6}, // "*b*b***b|**b*b***|***b***r|**B*****|********|********|*****r**|r***r***"
		{"b", 2, 3, 3, 4}, // "*b*b***b|**b*b***|***b***r|********|***B****|********|*****r**|r***r***"
		{"r", 0, 7, 1, 6}, // "*b*b***b|**b*b***|***b***r|********|***B****|********|*r***r**|****r***"
		{"b", 3, 2, 4, 3}, // "*b*b***b|**b*b***|*******r|****b***|***B****|********|*r***r**|****r***"
		{"r", 7, 2, 6, 1}, // "*b*b***b|**b*b*r*|********|****b***|***B****|********|*r***r**|****r***"
		{"b", 7, 0, 5, 2}, // "*b*b****|**b*b***|*****b**|****b***|***B****|********|*r***r**|****r***"
		{"r", 1, 6, 2, 5}, // "*b*b****|**b*b***|*****b**|****b***|***B****|**r*****|*****r**|****r***"
		{"b", 3, 4, 1, 6}, // "*b*b****|**b*b***|*****b**|****b***|********|********|*B***r**|****r***"
		{"r", 4, 7, 3, 6}, // "*b*b****|**b*b***|*****b**|****b***|********|********|*B*r*r**|********"
		{"b", 4, 3, 3, 4}, // "*b*b****|**b*b***|*****b**|********|***b****|********|*B*r*r**|********"
		{"r", 5, 6, 4, 5}, // "*b*b****|**b*b***|*****b**|********|***b****|****r***|*B*r****|********"
		{"b", 3, 4, 5, 6}, // "*b*b****|**b*b***|*****b**|********|********|********|*B*r*b**|********"
		{"r", 3, 6, 2, 5}, // "*b*b****|**b*b***|*****b**|********|********|**r*****|*B***b**|********"
		{"b", 1, 6, 3, 4}, // "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********"
	}
)

func GetPlayer(color string, black string, red string) string {
	if color == "b" {
		return black
	}
	return red
}

func PlayAllMoves(
	t *testing.T,
	msgServer types.MsgServer,
	context context.Context,
	gameIndex string,
	black string,
	red string,
	moves []GameMoveTest) {
	for _, move := range moves {
		_, err := msgServer.PlayMove(context, &types.MsgPlayMove{
			Creator:   GetPlayer(move.player, black, red),
			GameIndex: gameIndex,
			FromX:     move.fromX,
			FromY:     move.fromY,
			ToX:       move.toX,
			ToY:       move.toY,
		})
		require.Nil(t, err)
	}
}
