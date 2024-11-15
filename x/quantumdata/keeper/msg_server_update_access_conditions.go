package keeper

import (
	"context"

	"mdc-chain/x/quantumdata/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAccessConditions(goCtx context.Context, msg *types.MsgUpdateAccessConditions) (*types.MsgUpdateAccessConditionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateAccessConditionsResponse{}, nil
}
