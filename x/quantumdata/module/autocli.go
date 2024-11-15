package quantumdata

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "mdc-chain/api/mdcchain/quantumdata"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateQuantumDataFragment",
					Use:            "create-quantum-data-fragment [id] [data] [expiry] [access-level]",
					Short:          "Send a createQuantumDataFragment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "data"}, {ProtoField: "expiry"}, {ProtoField: "accessLevel"}},
				},
				{
					RpcMethod:      "AccessQuantumDataFragment",
					Use:            "access-quantum-data-fragment [id]",
					Short:          "Send a accessQuantumDataFragment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "UpdateAccessConditions",
					Use:            "update-access-conditions [id] [conditions]",
					Short:          "Send a updateAccessConditions tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "conditions"}},
				},
				{
					RpcMethod:      "UpdateQuantumDataStatus",
					Use:            "update-quantum-data-status [id] [new-status]",
					Short:          "Send a updateQuantumDataStatus tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "newStatus"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
