package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProperty{}, "realestate/CreateProperty", nil)
	cdc.RegisterConcrete(&MsgUpdateProperty{}, "realestate/UpdateProperty", nil)
	cdc.RegisterConcrete(&MsgDeleteProperty{}, "realestate/DeleteProperty", nil)
	cdc.RegisterConcrete(&MsgCreateCertificate{}, "realestate/CreateCertificate", nil)
	cdc.RegisterConcrete(&MsgUpdateCertificate{}, "realestate/UpdateCertificate", nil)
	cdc.RegisterConcrete(&MsgDeleteCertificate{}, "realestate/DeleteCertificate", nil)
	cdc.RegisterConcrete(&MsgCreateLocking{}, "realestate/CreateLocking", nil)
	cdc.RegisterConcrete(&MsgUpdateLocking{}, "realestate/UpdateLocking", nil)
	cdc.RegisterConcrete(&MsgDeleteLocking{}, "realestate/DeleteLocking", nil)
	cdc.RegisterConcrete(&MsgCreateInspection{}, "realestate/CreateInspection", nil)
	cdc.RegisterConcrete(&MsgUpdateInspection{}, "realestate/UpdateInspection", nil)
	cdc.RegisterConcrete(&MsgDeleteInspection{}, "realestate/DeleteInspection", nil)
	cdc.RegisterConcrete(&MsgCreateTransference{}, "realestate/CreateTransference", nil)
	cdc.RegisterConcrete(&MsgUpdateTransference{}, "realestate/UpdateTransference", nil)
	cdc.RegisterConcrete(&MsgDeleteTransference{}, "realestate/DeleteTransference", nil)
	cdc.RegisterConcrete(&MsgCreateHouse{}, "realestate/CreateHouse", nil)
	cdc.RegisterConcrete(&MsgUpdateHouse{}, "realestate/UpdateHouse", nil)
	cdc.RegisterConcrete(&MsgDeleteHouse{}, "realestate/DeleteHouse", nil)
	cdc.RegisterConcrete(&MsgRequestCertification{}, "realestate/RequestCertification", nil)
	cdc.RegisterConcrete(&MsgEmitCertificate{}, "realestate/EmitCertificate", nil)
	cdc.RegisterConcrete(&MsgLockProperty{}, "realestate/LockProperty", nil)
	cdc.RegisterConcrete(&MsgUnlockProperty{}, "realestate/UnlockProperty", nil)
	cdc.RegisterConcrete(&MsgUnlockAssets{}, "realestate/UnlockAssets", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProperty{},
		&MsgUpdateProperty{},
		&MsgDeleteProperty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCertificate{},
		&MsgUpdateCertificate{},
		&MsgDeleteCertificate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLocking{},
		&MsgUpdateLocking{},
		&MsgDeleteLocking{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInspection{},
		&MsgUpdateInspection{},
		&MsgDeleteInspection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTransference{},
		&MsgUpdateTransference{},
		&MsgDeleteTransference{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHouse{},
		&MsgUpdateHouse{},
		&MsgDeleteHouse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestCertification{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEmitCertificate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLockProperty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnlockProperty{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnlockAssets{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
