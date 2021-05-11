package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateOwner{}, "property/CreateOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateOwner{}, "property/UpdateOwner", nil)
	cdc.RegisterConcrete(&MsgDeleteOwner{}, "property/DeleteOwner", nil)

	cdc.RegisterConcrete(&MsgCreateProperty{}, "property/CreateProperty", nil)
	cdc.RegisterConcrete(&MsgUpdateProperty{}, "property/UpdateProperty", nil)
	cdc.RegisterConcrete(&MsgDeleteProperty{}, "property/DeleteProperty", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOwner{},
		&MsgUpdateOwner{},
		&MsgDeleteOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProperty{},
		&MsgUpdateProperty{},
		&MsgDeleteProperty{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
