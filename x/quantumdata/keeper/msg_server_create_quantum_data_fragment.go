package keeper

import (
	"context"

	"mdc-chain/x/quantumdata/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateQuantumDataFragment(goCtx context.Context, msg *types.MsgCreateQuantumDataFragment) (*types.MsgCreateQuantumDataFragmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateQuantumDataFragmentResponse{}, nil
}
