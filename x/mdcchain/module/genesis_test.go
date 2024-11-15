package mdcchain_test

import (
	"testing"

	keepertest "mdc-chain/testutil/keeper"
	"mdc-chain/testutil/nullify"
	mdcchain "mdc-chain/x/mdcchain/module"
	"mdc-chain/x/mdcchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MdcchainKeeper(t)
	mdcchain.InitGenesis(ctx, k, genesisState)
	got := mdcchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
