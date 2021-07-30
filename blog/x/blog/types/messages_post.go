// x/blog/types/messages_post.go
package types

import (
		sdk "github.com/cosmos/cosmos-sdk/types"
		sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

/*
A constructor function that creates the MsgCreatePost message
*/
func NewMsgCreatePost(creator string, title string, body string) *MsgCreatePost {
	return &MsgCreatePost{
				Creator: creator,
				Title: title,
				Body: body,
	}
}

/*
We have to define the below five functions
to implement the Msg interface. They allow
us to perform validation that doesn't require
access to the store(like checking for empty
values)
*/

// Route
func (msg MsgCreatePost) Route() string {
		return RouterKey
}

// Type ...
func (msg MsgCreatePost) Type() string {
		return "CreatePost"
}

// GetSigners ...
func (msg *MsgCreatePost) GetSigners() []sdk.AccAddress {
		creator, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
				panic(err)
		}
		return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *MsgCreatePost) GetSignBytes() []byte {
		bz := ModuleCdc.MustMarshalJSON(msg)
		return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePost) ValidateBasic() error {
		_, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
		}
		return nil
}