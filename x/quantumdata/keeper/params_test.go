package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "mdc-chain/testutil/keeper"
	"mdc-chain/x/quantumdata/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.QuantumdataKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
