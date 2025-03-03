package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) CommitCouncilResponse(goCtx context.Context, msg *types.MsgCommitCouncilResponse) (*types.MsgCommitCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collateralDeposit := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	council := k.Councils.Get(ctx, msg.CouncilId)
	if !slices.Contains(council.Voters, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	if council.Status != types.CouncelingStatus_councilCreated {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_councilCreated.String())
	}

	var allreadyVoted []string
	for _, response := range council.HashResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if slices.Contains(allreadyVoted, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Allready voted")
	}

	resp := types.WrapHashResponse{
        User: msg.Creator,
        Hash: msg.Response,
    }
	council.HashResponses = append(council.HashResponses, &resp)
	if msg.Suggestion != "" { // Direcly reveal when a suggestion is made
		clearResp := types.WrapClearResponse{
            User: msg.Creator,
            Response: types.Response_Suggestion,
            Suggestion: msg.Suggestion,
        }
		council.ClearResponses = append(council.ClearResponses, &clearResp)
	}

	if len(council.HashResponses) == 5 {
		council.Status = types.CouncelingStatus_commited
	}

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{collateralDeposit})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(collateralDeposit)

	err = k.TryEvaluate(ctx, council)
	if err != nil {
		return nil, err
	}

	k.Councils.Set(ctx, msg.CouncilId, council)

	return &types.MsgCommitCouncilResponseResponse{}, nil
}
