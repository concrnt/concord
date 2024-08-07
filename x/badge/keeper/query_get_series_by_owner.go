package keeper

import (
	"context"

	"cosmossdk.io/x/nft"
	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSeriesByOwner(goCtx context.Context, req *types.QueryGetSeriesByOwnerRequest) (*types.QueryGetSeriesByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	classes := k.nftKeeper.GetClasses(ctx)

	var series []*nft.Class
	for _, class := range classes {
		var classData types.ISeriesData
		err := k.cdc.UnpackAny(class.Data, &classData)
		if err != nil {
			return nil, err
		}

		if classData.GetCreator() == owner.String() {
			series = append(series, class)
		}
	}

	return &types.QueryGetSeriesByOwnerResponse{
		Series: series,
	}, nil
}
