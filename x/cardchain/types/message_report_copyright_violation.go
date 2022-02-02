package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReportCopyrightViolation = "report_copyright_violation"

var _ sdk.Msg = &MsgReportCopyrightViolation{}

func NewMsgReportCopyrightViolation(creator string, cardId uint64, originalArtist string, link string, additionalInfo string) *MsgReportCopyrightViolation {
	return &MsgReportCopyrightViolation{
		Creator:        creator,
		CardId:         cardId,
		OriginalArtist: originalArtist,
		Link:           link,
		AdditionalInfo: additionalInfo,
	}
}

func (msg *MsgReportCopyrightViolation) Route() string {
	return RouterKey
}

func (msg *MsgReportCopyrightViolation) Type() string {
	return TypeMsgReportCopyrightViolation
}

func (msg *MsgReportCopyrightViolation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReportCopyrightViolation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReportCopyrightViolation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
