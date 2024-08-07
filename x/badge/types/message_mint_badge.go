package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintBadge{}

func NewMsgMintBadge(creator string, series string, receiver string) *MsgMintBadge {
	return &MsgMintBadge{
		Creator:  creator,
		Series:   series,
		Receiver: receiver,
	}
}

func (msg *MsgMintBadge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
