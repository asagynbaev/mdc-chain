package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateQuantumDataFragment{}

func NewMsgCreateQuantumDataFragment(creator string, id string, data string, expiry string, accessLevel string) *MsgCreateQuantumDataFragment {
	return &MsgCreateQuantumDataFragment{
		Creator:     creator,
		Id:          id,
		Data:        data,
		Expiry:      expiry,
		AccessLevel: accessLevel,
	}
}

func (msg *MsgCreateQuantumDataFragment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
