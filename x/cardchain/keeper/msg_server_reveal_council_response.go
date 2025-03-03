package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) RevealCouncilResponse(goCtx context.Context, msg *types.MsgRevealCouncilResponse) (*types.MsgRevealCouncilResponseResponse, error) {
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

	if council.Status != types.CouncelingStatus_commited {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_commited.String())
	}

	var allreadyVoted []string
	for _, response := range council.ClearResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if slices.Contains(allreadyVoted, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already voted")
	}

	hashStringResponse := GetResponseHash(msg.Response, msg.Secret)
	var originalHash string

	for _, hash := range council.HashResponses {
		if hash.User == msg.Creator {
			originalHash = hash.Hash
			break
		}
	}

	if originalHash != hashStringResponse {
		return nil, types.ErrBadReveal
	}

	resp := types.WrapClearResponse{User: msg.Creator, Response: msg.Response}
	council.ClearResponses = append(council.ClearResponses, &resp)

	if len(council.ClearResponses) == 5 {
		council.Status = types.CouncelingStatus_revealed
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

	return &types.MsgRevealCouncilResponseResponse{}, nil
}
