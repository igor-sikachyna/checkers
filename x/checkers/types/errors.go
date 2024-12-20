package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidSigner           = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidBlack            = sdkerrors.Register(ModuleName, 1101, "black address is invalid: %s")
	ErrInvalidRed              = sdkerrors.Register(ModuleName, 1102, "red address is invalid: %s")
	ErrGameNotParseable        = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")
	ErrInvalidGameIndex        = sdkerrors.Register(ModuleName, 1104, "game index is invalid")
	ErrInvalidPositionIndex    = sdkerrors.Register(ModuleName, 1105, "position index is invalid")
	ErrMoveAbsent              = sdkerrors.Register(ModuleName, 1106, "there is no move")
	ErrGameNotFound            = sdkerrors.Register(ModuleName, 1107, "game by id not found")
	ErrCreatorNotPlayer        = sdkerrors.Register(ModuleName, 1108, "message creator is not a player")
	ErrNotPlayerTurn           = sdkerrors.Register(ModuleName, 1109, "player tried to play out of turn")
	ErrWrongMove               = sdkerrors.Register(ModuleName, 1110, "wrong move")
	ErrGameFinished            = sdkerrors.Register(ModuleName, 1111, "game is already finished")
	ErrInvalidDeadline         = sdkerrors.Register(ModuleName, 1112, "deadline cannot be parsed: %s")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1113, "cannot find winner by color: %s")
	ErrBlackCannotPay          = sdkerrors.Register(ModuleName, 1114, "black cannot pay the wager")
	ErrRedCannotPay            = sdkerrors.Register(ModuleName, 1115, "red cannot pay the wager")
	ErrNothingToPay            = sdkerrors.Register(ModuleName, 1116, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager       = sdkerrors.Register(ModuleName, 1117, "cannot refund wager to: %s")
	ErrCannotPayWinnings       = sdkerrors.Register(ModuleName, 1118, "cannot pay winnings to winner: %s")
	ErrNotInRefundState        = sdkerrors.Register(ModuleName, 1119, "game is not in a state to refund, move count: %d")
)
