package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/igor-sikachyna/checkers/x/checkers/rules"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, errorsmod.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, errorsmod.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, errorsmod.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, errorsmod.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetDeadlineAsTime()
	return err
}

func (storedGame StoredGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool, err error) {
	black, err := storedGame.GetBlackAddress()
	if err != nil {
		return nil, false, err
	}
	red, err := storedGame.GetRedAddress()
	if err != nil {
		return nil, false, err
	}
	address, found = map[string]sdk.AccAddress{
		rules.PieceStrings[rules.BLACK_PLAYER]: black,
		rules.PieceStrings[rules.RED_PLAYER]:   red,
	}[color]
	return address, found, nil
}

func (storedGame StoredGame) GetWinnerAddress() (address sdk.AccAddress, found bool, err error) {
	return storedGame.GetPlayerAddress(storedGame.Winner)
}

func (storedGame StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, errDeadline := time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, errorsmod.Wrapf(errDeadline, ErrInvalidDeadline.Error(), storedGame.Deadline)
}

func FormatDeadline(deadline time.Time) string {
	return deadline.UTC().Format(DeadlineLayout)
}

func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(MaxTurnDuration)
}

func (storedGame *StoredGame) GetWagerCoin() (wager sdk.Coin) {
	return sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(int64(storedGame.Wager)))
}
