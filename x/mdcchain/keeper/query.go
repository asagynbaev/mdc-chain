package keeper

import (
	"mdc-chain/x/mdcchain/types"
)

var _ types.QueryServer = Keeper{}
