package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/concrnt/concord/testutil/keeper"
	"github.com/concrnt/concord/x/badge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BadgeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
