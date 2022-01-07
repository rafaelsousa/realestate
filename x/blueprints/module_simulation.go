package blueprints

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/rafaelsousa/realestate/testutil/sample"
	blueprintssimulation "github.com/rafaelsousa/realestate/x/blueprints/simulation"
	"github.com/rafaelsousa/realestate/x/blueprints/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = blueprintssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHouse int = 100

	opWeightMsgUpdateHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHouse int = 100

	opWeightMsgDeleteHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHouse int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blueprintsGenesis := types.GenesisState{
		HouseList: []types.House{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		HouseCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blueprintsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHouse, &weightMsgCreateHouse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHouse = defaultWeightMsgCreateHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHouse,
		blueprintssimulation.SimulateMsgCreateHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHouse, &weightMsgUpdateHouse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHouse = defaultWeightMsgUpdateHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHouse,
		blueprintssimulation.SimulateMsgUpdateHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHouse, &weightMsgDeleteHouse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHouse = defaultWeightMsgDeleteHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHouse,
		blueprintssimulation.SimulateMsgDeleteHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
