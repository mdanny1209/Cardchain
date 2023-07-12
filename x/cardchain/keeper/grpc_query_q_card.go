package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCard(goCtx context.Context, req *types.QueryQCardRequest) (*types.OutpCard, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start of query code
	cardId, err := strconv.ParseUint(req.CardId, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not parse cardId")
	}

	card := k.Cards.Get(ctx, cardId)
	if card == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)
	sum := md5.Sum(image.Image)
	hash := hex.EncodeToString(sum[:])

	outpCard := types.OutpCard{
		Owner:              card.Owner,
		Artist:             card.Artist,
		Content:            string(card.Content),
		Image:              string(image.Image),
		FullArt:            card.FullArt,
		Notes:              card.Notes,
		Status:             card.Status,
		VotePool:           card.VotePool,
		Voters:             card.Voters,
		FairEnoughVotes:    card.FairEnoughVotes,
		OverpoweredVotes:   card.OverpoweredVotes,
		UnderpoweredVotes:  card.UnderpoweredVotes,
		InappropriateVotes: card.InappropriateVotes,
		Nerflevel:          card.Nerflevel,
<<<<<<< HEAD
		BalanceAnchor:			card.BalanceAnchor,
=======
		BalanceAnchor:      card.BalanceAnchor,
>>>>>>> d9d89c49e39a8b7d4657fa3fc10957282389e5bb
		Hash:               hash,
	}

	return &outpCard, nil
}
