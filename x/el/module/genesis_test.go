package el_test

import (
	"testing"

	keepertest "el/testutil/keeper"
	"el/testutil/nullify"
	el "el/x/el/module"
	"el/x/el/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ElKeeper(t)
	el.InitGenesis(ctx, k, genesisState)
	got := el.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
