package keeper

import (
	"context"

	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSeries(goCtx context.Context, msg *types.MsgCreateSeries) (*types.MsgCreateSeriesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateSeriesResponse{}, nil
}
