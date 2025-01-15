package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/igor-sikachyna/checkers/testutil/keeper"
	"github.com/igor-sikachyna/checkers/testutil/nullify"
	"github.com/igor-sikachyna/checkers/x/leaderboard/types"
)

func TestLeaderboardQuery(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetLeaderboardRequest
		response *types.QueryGetLeaderboardResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLeaderboardRequest{},
			response: &types.QueryGetLeaderboardResponse{Leaderboard: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Leaderboard(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
