package simulation

import (
	"math/rand"

	"mdc-chain/x/quantumdata/keeper"
	"mdc-chain/x/quantumdata/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateQuantumDataStatus(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateQuantumDataStatus{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateQuantumDataStatus simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateQuantumDataStatus simulation not implemented"), nil, nil
	}
}
