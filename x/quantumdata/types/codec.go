package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateQuantumDataFragment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccessQuantumDataFragment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateAccessConditions{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateQuantumDataStatus{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
