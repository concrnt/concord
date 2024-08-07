package keeper

import (
	"context"
	"crypto/md5"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/x/nft"
	"github.com/concrnt/concord/x/badge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/totegamma/concurrent/cdid"
)

func (k msgServer) MintBadge(goCtx context.Context, msg *types.MsgMintBadge) (*types.MsgMintBadgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	class, ok := k.nftKeeper.GetClass(ctx, msg.Series)
	if !ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "class not found")
	}
	var badgeData types.SeriesData
	err = badgeData.Unmarshal(class.Data.Value)
	if err != nil {
		return nil, err
	}

	if badgeData.Creator != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the creator can mint badges")
	}

	source := []byte(msg.String())
	hash := md5.Sum(source)
	var first10 [10]byte
	copy(first10[:], hash[:10])

	id := "b" + cdid.New(first10, ctx.BlockTime()).String()

	err = k.nftKeeper.Mint(ctx, nft.NFT{
		ClassId: msg.Series,
		Uri:     msg.Uri,
		Id:      id,
	}, receiver)

	if err != nil {
		return nil, err
	}

	return &types.MsgMintBadgeResponse{}, nil
}
