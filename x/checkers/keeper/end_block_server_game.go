package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/igor-sikachyna/checkers/x/checkers/rules"
	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	opponents := map[string]string{
		rules.PieceStrings[rules.BLACK_PLAYER]: rules.PieceStrings[rules.RED_PLAYER],
		rules.PieceStrings[rules.RED_PLAYER]:   rules.PieceStrings[rules.BLACK_PLAYER],
	}

	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	gameIndex := systemInfo.FifoHeadIndex
	var storedGame types.StoredGame

	for {
		if gameIndex == types.NoFifoIndex {
			break
		}

		storedGame, found = k.GetStoredGame(ctx, gameIndex)
		if !found {
			panic("Fifo head game not found " + systemInfo.FifoHeadIndex)
		}
		deadline, err := storedGame.GetDeadlineAsTime()
		if err != nil {
			panic(err)
		}

		if deadline.Before(ctx.BlockTime()) {
			k.RemoveFromFifo(ctx, &storedGame, &systemInfo)

			lastBoard := storedGame.Board
			if storedGame.MoveCount <= 1 {
				// No point in keeping a game that was never really played
				k.RemoveStoredGame(ctx, gameIndex)
				if storedGame.MoveCount == 1 {
					k.MustRefundWager(ctx, &storedGame)
				}
			} else {
				storedGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Turn))
				}
				k.MustPayWinnings(ctx, &storedGame)
				storedGame.Board = ""
				k.SetStoredGame(ctx, storedGame)
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.GameForfeitedEventType,
					sdk.NewAttribute(types.GameForfeitedEventGameIndex, gameIndex),
					sdk.NewAttribute(types.GameForfeitedEventWinner, storedGame.Winner),
					sdk.NewAttribute(types.GameForfeitedEventBoard, lastBoard),
				),
			)
		} else {
			// All other games after are active anyway
			break
		}

		gameIndex = systemInfo.FifoHeadIndex
	}

	k.SetSystemInfo(ctx, systemInfo)
}
