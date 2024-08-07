package badge_test

import (
	"testing"

	keepertest "github.com/concrnt/concord/testutil/keeper"
	"github.com/concrnt/concord/testutil/nullify"
	badge "github.com/concrnt/concord/x/badge/module"
	"github.com/concrnt/concord/x/badge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BadgeKeeper(t)
	badge.InitGenesis(ctx, k, genesisState)
	got := badge.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
