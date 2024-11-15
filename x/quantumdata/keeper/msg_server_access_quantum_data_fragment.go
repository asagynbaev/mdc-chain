package keeper

import (
	"context"

	"mdc-chain/x/quantumdata/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AccessQuantumDataFragment(goCtx context.Context, msg *types.MsgAccessQuantumDataFragment) (*types.MsgAccessQuantumDataFragmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAccessQuantumDataFragmentResponse{}, nil
}
