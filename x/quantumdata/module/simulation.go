package quantumdata

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"mdc-chain/testutil/sample"
	quantumdatasimulation "mdc-chain/x/quantumdata/simulation"
	"mdc-chain/x/quantumdata/types"
)

// avoid unused import issue
var (
	_ = quantumdatasimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateQuantumDataFragment = "op_weight_msg_create_quantum_data_fragment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateQuantumDataFragment int = 100

	opWeightMsgAccessQuantumDataFragment = "op_weight_msg_access_quantum_data_fragment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccessQuantumDataFragment int = 100

	opWeightMsgUpdateAccessConditions = "op_weight_msg_update_access_conditions"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccessConditions int = 100

	opWeightMsgUpdateQuantumDataStatus = "op_weight_msg_update_quantum_data_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateQuantumDataStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	quantumdataGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&quantumdataGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateQuantumDataFragment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateQuantumDataFragment, &weightMsgCreateQuantumDataFragment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateQuantumDataFragment = defaultWeightMsgCreateQuantumDataFragment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateQuantumDataFragment,
		quantumdatasimulation.SimulateMsgCreateQuantumDataFragment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccessQuantumDataFragment int
	simState.AppParams.GetOrGenerate(opWeightMsgAccessQuantumDataFragment, &weightMsgAccessQuantumDataFragment, nil,
		func(_ *rand.Rand) {
			weightMsgAccessQuantumDataFragment = defaultWeightMsgAccessQuantumDataFragment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccessQuantumDataFragment,
		quantumdatasimulation.SimulateMsgAccessQuantumDataFragment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAccessConditions int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAccessConditions, &weightMsgUpdateAccessConditions, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAccessConditions = defaultWeightMsgUpdateAccessConditions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAccessConditions,
		quantumdatasimulation.SimulateMsgUpdateAccessConditions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateQuantumDataStatus int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateQuantumDataStatus, &weightMsgUpdateQuantumDataStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateQuantumDataStatus = defaultWeightMsgUpdateQuantumDataStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateQuantumDataStatus,
		quantumdatasimulation.SimulateMsgUpdateQuantumDataStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateQuantumDataFragment,
			defaultWeightMsgCreateQuantumDataFragment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				quantumdatasimulation.SimulateMsgCreateQuantumDataFragment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAccessQuantumDataFragment,
			defaultWeightMsgAccessQuantumDataFragment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				quantumdatasimulation.SimulateMsgAccessQuantumDataFragment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAccessConditions,
			defaultWeightMsgUpdateAccessConditions,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				quantumdatasimulation.SimulateMsgUpdateAccessConditions(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateQuantumDataStatus,
			defaultWeightMsgUpdateQuantumDataStatus,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				quantumdatasimulation.SimulateMsgUpdateQuantumDataStatus(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
