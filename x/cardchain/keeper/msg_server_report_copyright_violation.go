package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ReportCopyrightViolation(goCtx context.Context, msg *types.MsgReportCopyrightViolation) (*types.MsgReportCopyrightViolationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReportCopyrightViolationResponse{}, nil
}
