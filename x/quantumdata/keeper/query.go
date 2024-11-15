package keeper

import (
	"mdc-chain/x/quantumdata/types"
)

var _ types.QueryServer = Keeper{}
