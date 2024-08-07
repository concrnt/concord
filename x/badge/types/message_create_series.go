package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSeries{}

func NewMsgCreateSeries(creator string, name string, description string, uri string, transferable bool) *MsgCreateSeries {
	return &MsgCreateSeries{
		Creator:      creator,
		Name:         name,
		Description:  description,
		Uri:          uri,
		Transferable: transferable,
	}
}

func (msg *MsgCreateSeries) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
