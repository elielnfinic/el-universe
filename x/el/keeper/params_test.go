package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "el/testutil/keeper"
	"el/x/el/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ElKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
