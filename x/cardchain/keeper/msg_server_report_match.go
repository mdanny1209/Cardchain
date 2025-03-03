package keeper

import (
	"context"
	"time"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportMatch(goCtx context.Context, msg *types.MsgReportMatch) (*types.MsgReportMatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}

	matchId := k.Matches.GetNum(ctx)

	match := types.Match{
		Timestamp:        uint64(time.Now().Unix()),
		Reporter:         msg.Creator,
		PlayerA:          types.NewMatchPlayer(msg.PlayerA, msg.CardsA),
		PlayerB:          types.NewMatchPlayer(msg.PlayerB, msg.CardsB),
		Outcome:          msg.Outcome,
		CoinsDistributed: false,
	}

	_, err = k.GetMatchAddresses(ctx, match)
	if err != nil {
		return nil, err
	}

	if msg.Outcome == types.Outcome_Aborted {
		err = k.DistributeCoins(ctx, &match, msg.Outcome)
        if err != nil {
            return nil, err
        }
		k.ReportServerMatch(ctx, msg.Creator, 1, true)
	}

	k.Matches.Set(ctx, matchId, &match)

	return &types.MsgReportMatchResponse{}, nil
}
