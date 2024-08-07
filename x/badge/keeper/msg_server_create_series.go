package keeper

import (
	"context"
	"crypto/md5"

	"cosmossdk.io/x/nft"
	"github.com/concrnt/concord/x/badge/types"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/totegamma/concurrent/cdid"
)

func (k msgServer) CreateSeries(goCtx context.Context, msg *types.MsgCreateSeries) (*types.MsgCreateSeriesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	source := []byte(msg.String())
	hash := md5.Sum(source)
	var first10 [10]byte
	copy(first10[:], hash[:10])

	id := "B" + cdid.New(first10, ctx.BlockTime()).String()

	badgeData := types.SeriesData{
		Creator:      msg.Creator,
		Transferable: msg.Transferable,
	}

	anydata, err := codec.NewAnyWithValue(&badgeData)
	if err != nil {
		return nil, err
	}

	k.nftKeeper.SaveClass(ctx, nft.Class{
		Id:          id,
		Name:        msg.Name,
		Description: msg.Description,
		Uri:         msg.Uri,
		Data:        anydata,
	})

	return &types.MsgCreateSeriesResponse{}, nil
}
