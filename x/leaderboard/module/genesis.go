package leaderboard

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/igor-sikachyna/checkers/x/leaderboard/keeper"
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetLeaderboard(ctx, genState.Leaderboard)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all leaderboard
	leaderboard, found := k.GetLeaderboard(ctx)
	if found {
		genesis.Leaderboard = leaderboard
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
