package keeper

import (
	"github.com/igor-sikachyna/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
