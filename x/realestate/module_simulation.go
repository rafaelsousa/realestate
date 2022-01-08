package realestate

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/rafaelsousa/realestate/testutil/sample"
	realestatesimulation "github.com/rafaelsousa/realestate/x/realestate/simulation"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = realestatesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProperty = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProperty int = 100

	opWeightMsgUpdateProperty = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProperty int = 100

	opWeightMsgDeleteProperty = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProperty int = 100

	opWeightMsgCreateCertificate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCertificate int = 100

	opWeightMsgUpdateCertificate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCertificate int = 100

	opWeightMsgDeleteCertificate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCertificate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	realestateGenesis := types.GenesisState{
		PropertyList: []types.Property{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PropertyCount: 2,
			CertificateList: []types.Certificate{
		{
			Id: 0,
			Creator: sample.AccAddress(),

		},
		{
			Id: 1,
			Creator: sample.AccAddress(),

		},
	},
	CertificateCount: 2,
	// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&realestateGenesis)
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

	var weightMsgCreateProperty int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProperty, &weightMsgCreateProperty, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProperty = defaultWeightMsgCreateProperty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProperty,
		realestatesimulation.SimulateMsgCreateProperty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProperty int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProperty, &weightMsgUpdateProperty, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProperty = defaultWeightMsgUpdateProperty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProperty,
		realestatesimulation.SimulateMsgUpdateProperty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProperty int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProperty, &weightMsgDeleteProperty, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProperty = defaultWeightMsgDeleteProperty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProperty,
		realestatesimulation.SimulateMsgDeleteProperty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCertificate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCertificate, &weightMsgCreateCertificate, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCertificate = defaultWeightMsgCreateCertificate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCertificate,
		realestatesimulation.SimulateMsgCreateCertificate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCertificate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCertificate, &weightMsgUpdateCertificate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCertificate = defaultWeightMsgUpdateCertificate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCertificate,
		realestatesimulation.SimulateMsgUpdateCertificate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCertificate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCertificate, &weightMsgDeleteCertificate, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCertificate = defaultWeightMsgDeleteCertificate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCertificate,
		realestatesimulation.SimulateMsgDeleteCertificate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
