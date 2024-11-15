package types

import (
	"testing"

	"mdc-chain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateQuantumDataFragment_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateQuantumDataFragment
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateQuantumDataFragment{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateQuantumDataFragment{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
