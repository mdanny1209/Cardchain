package keeper

import (
	"context"
	"encoding/json"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SaveCardContent(goCtx context.Context, msg *types.MsgSaveCardContent) (*types.MsgSaveCardContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.Cards.Get(ctx, msg.CardId)

	msgOwner, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	// TODO
	//if card.Status != types.Status_prototype {
	//	return nil, sdkerrors.Wrap(types.ErrInvalidCardStatus, "Card has to be a prototype to be changeable")
	//}

	if card.Owner == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	cardobj, err := keywords.Unmarshal([]byte(msg.Content))
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())

	}

	card.Content, err = json.Marshal(cardobj)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCardobject, err.Error())
	}

	card.Notes = msg.Notes
	card.Artist = msg.Artist
	if card.Status == types.Status_scheme {
		err = k.TransferSchemeToCard(ctx, msg.CardId, &msgOwner)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "An error accured while converting a card to a scheme: "+err.Error())
		}
	}

	claimedAirdrop := k.ClaimAirDrop(ctx, &msgOwner, types.AirDrop_create)
	k.SetUserFromUser(ctx, msgOwner)

	// card.Status = types.Status_prototype
	card.Status = types.Status_permanent // TODO: remove later
	k.Cards.Set(ctx, msg.CardId, card)

	return &types.MsgSaveCardContentResponse{AirdropClaimed: claimedAirdrop}, nil
}
