package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

type ISeriesData interface {
	GetCreator() string
	GetTransferable() bool
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSeries{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintBadge{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*ISeriesData)(nil),
		&SeriesData{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
