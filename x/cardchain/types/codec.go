package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateuser{}, "cardchain/Createuser", nil)
	cdc.RegisterConcrete(&MsgBuyCardScheme{}, "cardchain/BuyCardScheme", nil)
	cdc.RegisterConcrete(&MsgVoteCard{}, "cardchain/VoteCard", nil)
	cdc.RegisterConcrete(&MsgSaveCardContent{}, "cardchain/SaveCardContent", nil)
	cdc.RegisterConcrete(&MsgTransferCard{}, "cardchain/TransferCard", nil)
	cdc.RegisterConcrete(&MsgDonateToCard{}, "cardchain/DonateToCard", nil)
	cdc.RegisterConcrete(&MsgAddArtwork{}, "cardchain/AddArtwork", nil)
	cdc.RegisterConcrete(&MsgSubmitCopyrightProposal{}, "cardchain/SubmitCopyrightProposal", nil)
	cdc.RegisterConcrete(&CopyrightProposal{}, "cardchain/CopyrightProposal", nil)
	cdc.RegisterConcrete(&MatchReporterProposal{}, "cardchain/MatchReporterProposal", nil)
	cdc.RegisterConcrete(&CollectionProposal{}, "cardchain/CollectionProposal", nil)
	cdc.RegisterConcrete(&MsgChangeArtist{}, "cardchain/ChangeArtist", nil)
	cdc.RegisterConcrete(&MsgRegisterForCouncil{}, "cardchain/RegisterForCouncil", nil)
	cdc.RegisterConcrete(&MsgReportMatch{}, "cardchain/ReportMatch", nil)
	cdc.RegisterConcrete(&MsgSubmitMatchReporterProposal{}, "cardchain/SubmitMatchReporterProposal", nil)
	cdc.RegisterConcrete(&MsgApointMatchReporter{}, "cardchain/ApointMatchReporter", nil)
	cdc.RegisterConcrete(&MsgCreateCollection{}, "cardchain/CreateCollection", nil)
	cdc.RegisterConcrete(&MsgAddCardToCollection{}, "cardchain/AddCardToCollection", nil)
	cdc.RegisterConcrete(&MsgFinalizeCollection{}, "cardchain/FinalizeCollection", nil)
	cdc.RegisterConcrete(&MsgBuyCollection{}, "cardchain/BuyCollection", nil)
	cdc.RegisterConcrete(&MsgRemoveCardFromCollection{}, "cardchain/RemoveCardFromCollection", nil)
	cdc.RegisterConcrete(&MsgRemoveContributorFromCollection{}, "cardchain/RemoveContributorFromCollection", nil)
	cdc.RegisterConcrete(&MsgAddContributorToCollection{}, "cardchain/AddContributorToCollection", nil)
	cdc.RegisterConcrete(&MsgSubmitCollectionProposal{}, "cardchain/SubmitCollectionProposal", nil)
	cdc.RegisterConcrete(&MsgCreateSellOffer{}, "cardchain/CreateSellOffer", nil)
	cdc.RegisterConcrete(&MsgBuyCard{}, "cardchain/BuyCard", nil)
	cdc.RegisterConcrete(&MsgRemoveSellOffer{}, "cardchain/RemoveSellOffer", nil)
	cdc.RegisterConcrete(&MsgAddArtworkToCollection{}, "cardchain/AddArtworkToCollection", nil)
	cdc.RegisterConcrete(&MsgAddStoryToCollection{}, "cardchain/AddStoryToCollection", nil)
	cdc.RegisterConcrete(&MsgSetCardRarity{}, "cardchain/SetCardRarity", nil)
	cdc.RegisterConcrete(&MsgCreateCouncil{}, "cardchain/CreateCouncil", nil)
	cdc.RegisterConcrete(&MsgCommitCouncilResponse{}, "cardchain/CommitCouncilResponse", nil)
	cdc.RegisterConcrete(&MsgRevealCouncilResponse{}, "cardchain/RevealCouncilResponse", nil)
	cdc.RegisterConcrete(&MsgRestartCouncil{}, "cardchain/RestartCouncil", nil)
	cdc.RegisterConcrete(&MsgRewokeCouncilRegistration{}, "cardchain/RewokeCouncilRegistration", nil)
	cdc.RegisterConcrete(&MsgConfirmMatch{}, "cardchain/ConfirmMatch", nil)
	cdc.RegisterConcrete(&MsgSetProfileCard{}, "cardchain/SetProfileCard", nil)
	cdc.RegisterConcrete(&MsgOpenBoosterPack{}, "cardchain/OpenBoosterPack", nil)
	cdc.RegisterConcrete(&MsgTransferBoosterPack{}, "cardchain/TransferBoosterPack", nil)
	cdc.RegisterConcrete(&MsgSetCollectionStoryWriter{}, "cardchain/SetCollectionStoryWriter", nil)
	cdc.RegisterConcrete(&MsgSetCollectionArtist{}, "cardchain/SetCollectionArtist", nil)
	cdc.RegisterConcrete(&MsgSetUserWebsite{}, "cardchain/SetUserWebsite", nil)
	cdc.RegisterConcrete(&MsgSetUserBiography{}, "cardchain/SetUserBiography", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateuser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCardScheme{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveCardContent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDonateToCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddArtwork{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitCopyrightProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CopyrightProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MatchReporterProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CollectionProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeArtist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterForCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReportMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitMatchReporterProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApointMatchReporter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddCardToCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalizeCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveCardFromCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveContributorFromCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddContributorToCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitCollectionProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSellOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveSellOffer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddArtworkToCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddStoryToCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCardRarity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitCouncilResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealCouncilResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRestartCouncil{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRewokeCouncilRegistration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmMatch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProfileCard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenBoosterPack{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferBoosterPack{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCollectionStoryWriter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCollectionArtist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetUserWebsite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetUserBiography{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
