package badge

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/concrnt/concord/api/concord/badge"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateSeries",
					Use:            "create-series [name] [description] [uri] [transferable]",
					Short:          "Send a create-series tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "description"}, {ProtoField: "uri"}, {ProtoField: "transferable"}},
				},
				{
					RpcMethod:      "MintBadge",
					Use:            "mint-badge [series] [receiver]",
					Short:          "Send a mint-badge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "series"}, {ProtoField: "receiver"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
