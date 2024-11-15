package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateQuantumDataStatus{}

func NewMsgUpdateQuantumDataStatus(creator string, id string, newStatus string) *MsgUpdateQuantumDataStatus {
	return &MsgUpdateQuantumDataStatus{
		Creator:   creator,
		Id:        id,
		NewStatus: newStatus,
	}
}

func (msg *MsgUpdateQuantumDataStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
