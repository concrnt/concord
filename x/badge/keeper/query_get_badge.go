package keeper

import (
	"context"

	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBadge(goCtx context.Context, req *types.QueryGetBadgeRequest) (*types.QueryGetBadgeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	class, ok := k.nftKeeper.GetClass(ctx, req.SeriesId)
	if !ok {
		return nil, status.Error(codes.NotFound, "class not found")
	}

	var classData types.ISeriesData
	err := k.cdc.UnpackAny(class.Data, &classData)
	if err != nil {
		return nil, err
	}

	nft, ok := k.nftKeeper.GetNFT(ctx, req.SeriesId, req.BadgeId)
	if !ok {
		return nil, status.Error(codes.NotFound, "nft not found")
	}

	uri := nft.Uri
	if uri == "" {
		uri = class.Uri
	}

	owner := k.nftKeeper.GetOwner(ctx, req.SeriesId, req.BadgeId)
	if owner == nil {
		return nil, status.Error(codes.NotFound, "owner not found")
	}

	badge := types.Badge{
		Name:         class.Name,
		Description:  class.Description,
		Uri:          uri,
		Creator:      classData.GetCreator(),
		Owner:        owner.String(),
		Transferable: classData.GetTransferable(),
	}

	return &types.QueryGetBadgeResponse{
		Badge: &badge,
	}, nil
}
