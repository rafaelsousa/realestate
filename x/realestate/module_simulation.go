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

	opWeightMsgCreateLocking = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLocking int = 100

	opWeightMsgUpdateLocking = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLocking int = 100

	opWeightMsgDeleteLocking = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLocking int = 100

	opWeightMsgCreateInspection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateInspection int = 100

	opWeightMsgUpdateInspection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateInspection int = 100

	opWeightMsgDeleteInspection = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteInspection int = 100

	opWeightMsgCreateTransference = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTransference int = 100

	opWeightMsgUpdateTransference = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTransference int = 100

	opWeightMsgDeleteTransference = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTransference int = 100

	opWeightMsgCreateHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateHouse int = 100

	opWeightMsgUpdateHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateHouse int = 100

	opWeightMsgDeleteHouse = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteHouse int = 100

	opWeightMsgRequestCertification = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestCertification int = 100

	opWeightMsgEmitCertificate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEmitCertificate int = 100

	opWeightMsgLockProperty = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLockProperty int = 100

	opWeightMsgUnlockProperty = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnlockProperty int = 100

	opWeightMsgUnlockAssets = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnlockAssets int = 100

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
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CertificateCount: 2,
		LockingList: []types.Locking{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		LockingCount: 2,
		InspectionList: []types.Inspection{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		InspectionCount: 2,
		TransferenceList: []types.Transference{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TransferenceCount: 2,
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

	var weightMsgCreateLocking int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLocking, &weightMsgCreateLocking, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLocking = defaultWeightMsgCreateLocking
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLocking,
		realestatesimulation.SimulateMsgCreateLocking(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLocking int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLocking, &weightMsgUpdateLocking, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLocking = defaultWeightMsgUpdateLocking
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLocking,
		realestatesimulation.SimulateMsgUpdateLocking(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLocking int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteLocking, &weightMsgDeleteLocking, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLocking = defaultWeightMsgDeleteLocking
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLocking,
		realestatesimulation.SimulateMsgDeleteLocking(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateInspection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateInspection, &weightMsgCreateInspection, nil,
		func(_ *rand.Rand) {
			weightMsgCreateInspection = defaultWeightMsgCreateInspection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateInspection,
		realestatesimulation.SimulateMsgCreateInspection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateInspection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateInspection, &weightMsgUpdateInspection, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateInspection = defaultWeightMsgUpdateInspection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateInspection,
		realestatesimulation.SimulateMsgUpdateInspection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteInspection int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteInspection, &weightMsgDeleteInspection, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteInspection = defaultWeightMsgDeleteInspection
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteInspection,
		realestatesimulation.SimulateMsgDeleteInspection(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTransference int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTransference, &weightMsgCreateTransference, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTransference = defaultWeightMsgCreateTransference
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTransference,
		realestatesimulation.SimulateMsgCreateTransference(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTransference int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTransference, &weightMsgUpdateTransference, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTransference = defaultWeightMsgUpdateTransference
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTransference,
		realestatesimulation.SimulateMsgUpdateTransference(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTransference int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTransference, &weightMsgDeleteTransference, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTransference = defaultWeightMsgDeleteTransference
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTransference,
		realestatesimulation.SimulateMsgDeleteTransference(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateHouse, &weightMsgCreateHouse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateHouse = defaultWeightMsgCreateHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateHouse,
		realestatesimulation.SimulateMsgCreateHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateHouse, &weightMsgUpdateHouse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateHouse = defaultWeightMsgUpdateHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateHouse,
		realestatesimulation.SimulateMsgUpdateHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteHouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteHouse, &weightMsgDeleteHouse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteHouse = defaultWeightMsgDeleteHouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteHouse,
		realestatesimulation.SimulateMsgDeleteHouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRequestCertification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestCertification, &weightMsgRequestCertification, nil,
		func(_ *rand.Rand) {
			weightMsgRequestCertification = defaultWeightMsgRequestCertification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestCertification,
		realestatesimulation.SimulateMsgRequestCertification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEmitCertificate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEmitCertificate, &weightMsgEmitCertificate, nil,
		func(_ *rand.Rand) {
			weightMsgEmitCertificate = defaultWeightMsgEmitCertificate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEmitCertificate,
		realestatesimulation.SimulateMsgEmitCertificate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLockProperty int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLockProperty, &weightMsgLockProperty, nil,
		func(_ *rand.Rand) {
			weightMsgLockProperty = defaultWeightMsgLockProperty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLockProperty,
		realestatesimulation.SimulateMsgLockProperty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnlockProperty int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnlockProperty, &weightMsgUnlockProperty, nil,
		func(_ *rand.Rand) {
			weightMsgUnlockProperty = defaultWeightMsgUnlockProperty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnlockProperty,
		realestatesimulation.SimulateMsgUnlockProperty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnlockAssets int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnlockAssets, &weightMsgUnlockAssets, nil,
		func(_ *rand.Rand) {
			weightMsgUnlockAssets = defaultWeightMsgUnlockAssets
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnlockAssets,
		realestatesimulation.SimulateMsgUnlockAssets(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
