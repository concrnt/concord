package keeper

import (
	"context"

	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBadgesBySeries(goCtx context.Context, req *types.QueryGetBadgesBySeriesRequest) (*types.QueryGetBadgesBySeriesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	class, ok := k.nftKeeper.GetClass(ctx, req.Series)
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}

	var classData types.ISeriesData
	err := k.cdc.UnpackAny(class.Data, &classData)
	if err != nil {
		return nil, err
	}

	nfts := k.nftKeeper.GetNFTsOfClass(ctx, req.Series)

	badges := make([]*types.Badge, len(nfts))
	for i, nft := range nfts {
		uri := nft.Uri
		if uri == "" {
			uri = class.Uri
		}

		owner := k.nftKeeper.GetOwner(ctx, class.Id, nft.Id)

		badge := types.Badge{
			ClassId:      class.Id,
			BadgeId:      nft.Id,
			Name:         class.Name,
			Description:  class.Description,
			Uri:          uri,
			Creator:      classData.GetCreator(),
			Owner:        owner.String(),
			Transferable: classData.GetTransferable(),
		}

		badges[i] = &badge
	}

	return &types.QueryGetBadgesBySeriesResponse{
		Badges: badges,
	}, nil
}
