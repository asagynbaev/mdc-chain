package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAccessConditions{}

func NewMsgUpdateAccessConditions(creator string, id string, conditions string) *MsgUpdateAccessConditions {
	return &MsgUpdateAccessConditions{
		Creator:    creator,
		Id:         id,
		Conditions: conditions,
	}
}

func (msg *MsgUpdateAccessConditions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
