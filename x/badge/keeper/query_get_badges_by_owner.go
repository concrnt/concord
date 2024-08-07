package keeper

import (
	"context"

	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBadgesByOwner(goCtx context.Context, req *types.QueryGetBadgesByOwnerRequest) (*types.QueryGetBadgesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	classes := k.nftKeeper.GetClasses(ctx)

	badges := make([]*types.Badge, 0)

	for _, class := range classes {
		nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, class.Id, owner)
		if len(nfts) < 0 {
			continue
		}

		var classData types.ISeriesData
		err := k.cdc.UnpackAny(class.Data, &classData)
		if err != nil {
			return nil, err
		}

		for _, nft := range nfts {
			uri := nft.Uri
			if uri == "" {
				uri = class.Uri
			}

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

			badges = append(badges, &badge)
		}
	}

	return &types.QueryGetBadgesByOwnerResponse{
		Badges: badges,
	}, nil
}
