package keeper

import (
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
