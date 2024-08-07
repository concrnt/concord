package badge

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/concrnt/concord/testutil/sample"
	badgesimulation "github.com/concrnt/concord/x/badge/simulation"
	"github.com/concrnt/concord/x/badge/types"
)

// avoid unused import issue
var (
	_ = badgesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateSeries = "op_weight_msg_create_series"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSeries int = 100

	opWeightMsgMintBadge = "op_weight_msg_mint_badge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintBadge int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	badgeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&badgeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSeries int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSeries, &weightMsgCreateSeries, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSeries = defaultWeightMsgCreateSeries
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSeries,
		badgesimulation.SimulateMsgCreateSeries(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintBadge int
	simState.AppParams.GetOrGenerate(opWeightMsgMintBadge, &weightMsgMintBadge, nil,
		func(_ *rand.Rand) {
			weightMsgMintBadge = defaultWeightMsgMintBadge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintBadge,
		badgesimulation.SimulateMsgMintBadge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSeries,
			defaultWeightMsgCreateSeries,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				badgesimulation.SimulateMsgCreateSeries(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintBadge,
			defaultWeightMsgMintBadge,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				badgesimulation.SimulateMsgMintBadge(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
