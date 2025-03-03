package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) AddCardToCollection(goCtx context.Context, msg *types.MsgAddCardToCollection) (*types.MsgAddCardToCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collectionSize := int(k.GetParams(ctx).CollectionSize)

	collection := k.Collections.Get(ctx, msg.CollectionId)
	if !slices.Contains(collection.Contributors, msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid contributor")
	}
	if collection.Status != types.CStatus_design {
		return nil, types.ErrCollectionNotInDesign
	}

	iter := k.Collections.GetItemIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		idx, coll := iter.Value()
		if coll.Status != types.CStatus_archived && slices.Contains(coll.Cards, msg.CardId) {
			return nil, sdkerrors.Wrapf(types.ErrCardAlreadyInCollection, "Collection: %d", idx)
		}
	}

	card := k.Cards.Get(ctx, msg.CardId)
	if card.Status != types.Status_permanent {
		return nil, sdkerrors.Wrap(types.ErrCardDoesNotExist, "Card is not permanent or does not exist")
	}

	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid creator")
	}

	if len(collection.Cards) >= collectionSize {
		return nil, sdkerrors.Wrapf(types.ErrCollectionSize, "Max is %d", collectionSize)
	}

	if slices.Contains(collection.Cards, msg.CardId) {
		return nil, sdkerrors.Wrapf(types.ErrCardAlreadyInCollection, "Card: %d", msg.CardId)
	}

	err := k.CollectCollectionConributionFee(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	collection.Cards = append(collection.Cards, msg.CardId)

	k.Collections.Set(ctx, msg.CollectionId, collection)

	return &types.MsgAddCardToCollectionResponse{}, nil
}
