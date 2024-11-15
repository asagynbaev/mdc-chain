package simulation

import (
	"math/rand"

	"mdc-chain/x/quantumdata/keeper"
	"mdc-chain/x/quantumdata/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateQuantumDataFragment(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateQuantumDataFragment{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateQuantumDataFragment simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "CreateQuantumDataFragment simulation not implemented"), nil, nil
	}
}