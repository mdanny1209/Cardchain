package keeper

import (
	"context"
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) VoteCard(goCtx context.Context, msg *types.MsgVoteCard) (*types.MsgVoteCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voter, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	voteRights := voter.VoteRights
	rightsIndex := slices.IndexFunc(voteRights, func(s *types.VoteRight) bool { return s.CardId == msg.CardId })

	// check if voting rights are true
	if rightsIndex < 0 {
		return nil, sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
	}

	//check if voting rights are timed out
	if ctx.BlockHeight() > voteRights[rightsIndex].ExpireBlock {
		k.RemoveVoteRight(ctx, &voter, rightsIndex)
		return nil, sdkerrors.Wrap(types.ErrVoteRightIsExpired, "Voting Right has expired")
	}

	// if the vote right is valid, get the Card
	card := k.Cards.Get(ctx, msg.CardId)

	// check if card status is valid
	if card.Status != types.Status_permanent && card.Status != types.Status_trial {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
	}

	switch msg.VoteType {
	case "fair_enough":
		card.FairEnoughVotes++
	case "inappropriate":
		card.InappropriateVotes++
	case "overpowered":
		card.OverpoweredVotes++
	case "underpowered":
		card.UnderpoweredVotes++
	default:
		errMsg := fmt.Sprintf("Unrecognized card vote type: %s", msg.VoteType)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
	card.Voters = append(card.Voters, msg.Creator)

	// check for specific bounty on the card
	if !card.VotePool.IsZero() {
        reward := k.GetParams(ctx).TrialVoteReward
		err := k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{reward})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		card.VotePool.Sub(reward) // TODO actually if there is less than 1cr then it should be adjusted
	}

	amount := k.GetVoteReward(ctx)
	k.MintCoinsToAddr(ctx, voter.Addr, sdk.Coins{amount})
	k.SubPoolCredits(ctx, BalancersPoolKey, amount)

	k.Cards.Set(ctx, msg.CardId, card)

	votes := k.RunningAverages.Get(ctx, Votes24ValueKey)
	votes.Arr[len(votes.Arr)-1]++
	k.RunningAverages.Set(ctx, Votes24ValueKey, votes)

	k.RemoveVoteRight(ctx, &voter, rightsIndex)

	claimedAirdrop := k.ClaimAirDrop(ctx, &voter, types.AirDrop_vote)
	k.SetUserFromUser(ctx, voter)

	return &types.MsgVoteCardResponse{AirdropClaimed: claimedAirdrop}, nil
}
