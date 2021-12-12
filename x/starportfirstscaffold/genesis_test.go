package starportfirstscaffold_test

import (
	"testing"

	keepertest "github.com/mkXultra/starport_first_scaffold/testutil/keeper"
	"github.com/mkXultra/starport_first_scaffold/x/starportfirstscaffold"
	"github.com/mkXultra/starport_first_scaffold/x/starportfirstscaffold/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StarportfirstscaffoldKeeper(t)
	starportfirstscaffold.InitGenesis(ctx, *k, genesisState)
	got := starportfirstscaffold.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
