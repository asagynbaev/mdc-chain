package quantumdata_test

import (
	"testing"

	keepertest "mdc-chain/testutil/keeper"
	"mdc-chain/testutil/nullify"
	quantumdata "mdc-chain/x/quantumdata/module"
	"mdc-chain/x/quantumdata/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QuantumdataKeeper(t)
	quantumdata.InitGenesis(ctx, k, genesisState)
	got := quantumdata.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
