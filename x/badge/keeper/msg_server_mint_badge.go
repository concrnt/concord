package keeper

import (
	"context"

	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintBadge(goCtx context.Context, msg *types.MsgMintBadge) (*types.MsgMintBadgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintBadgeResponse{}, nil
}
