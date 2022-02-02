import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export interface MsgCreateuser {
    creator: string;
    newUser: string;
    alias: string;
}
export interface MsgCreateuserResponse {
}
export interface MsgBuyCardScheme {
    creator: string;
    bid: string;
    buyer: string;
}
export interface MsgBuyCardSchemeResponse {
}
export interface MsgVoteCard {
    creator: string;
    cardId: number;
    voteType: string;
    voter: string;
}
export interface MsgVoteCardResponse {
}
export interface MsgSaveCardContent {
    creator: string;
    cardId: number;
    content: Uint8Array;
    /**
     * bytes image = 4;
     *  string fullArt = 5;
     */
    notes: string;
    owner: string;
    artist: string;
}
export interface MsgSaveCardContentResponse {
}
export interface MsgTransferCard {
    creator: string;
    cardId: number;
    sender: string;
    receiver: string;
}
export interface MsgTransferCardResponse {
}
export interface MsgDonateToCard {
    creator: string;
    cardId: number;
    donator: string;
    amount: string;
}
export interface MsgDonateToCardResponse {
}
export interface MsgAddArtwork {
    creator: string;
    cardId: number;
    image: Uint8Array;
    fullArt: boolean;
}
export interface MsgAddArtworkResponse {
}
export interface MsgReportCopyrightViolation {
    creator: string;
    cardId: number;
    originalArtist: string;
    link: string;
    additionalInfo: string;
}
export interface MsgReportCopyrightViolationResponse {
}
export declare const MsgCreateuser: {
    encode(message: MsgCreateuser, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuser;
    fromJSON(object: any): MsgCreateuser;
    toJSON(message: MsgCreateuser): unknown;
    fromPartial(object: DeepPartial<MsgCreateuser>): MsgCreateuser;
};
export declare const MsgCreateuserResponse: {
    encode(_: MsgCreateuserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateuserResponse;
    fromJSON(_: any): MsgCreateuserResponse;
    toJSON(_: MsgCreateuserResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateuserResponse>): MsgCreateuserResponse;
};
export declare const MsgBuyCardScheme: {
    encode(message: MsgBuyCardScheme, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardScheme;
    fromJSON(object: any): MsgBuyCardScheme;
    toJSON(message: MsgBuyCardScheme): unknown;
    fromPartial(object: DeepPartial<MsgBuyCardScheme>): MsgBuyCardScheme;
};
export declare const MsgBuyCardSchemeResponse: {
    encode(_: MsgBuyCardSchemeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBuyCardSchemeResponse;
    fromJSON(_: any): MsgBuyCardSchemeResponse;
    toJSON(_: MsgBuyCardSchemeResponse): unknown;
    fromPartial(_: DeepPartial<MsgBuyCardSchemeResponse>): MsgBuyCardSchemeResponse;
};
export declare const MsgVoteCard: {
    encode(message: MsgVoteCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVoteCard;
    fromJSON(object: any): MsgVoteCard;
    toJSON(message: MsgVoteCard): unknown;
    fromPartial(object: DeepPartial<MsgVoteCard>): MsgVoteCard;
};
export declare const MsgVoteCardResponse: {
    encode(_: MsgVoteCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVoteCardResponse;
    fromJSON(_: any): MsgVoteCardResponse;
    toJSON(_: MsgVoteCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgVoteCardResponse>): MsgVoteCardResponse;
};
export declare const MsgSaveCardContent: {
    encode(message: MsgSaveCardContent, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveCardContent;
    fromJSON(object: any): MsgSaveCardContent;
    toJSON(message: MsgSaveCardContent): unknown;
    fromPartial(object: DeepPartial<MsgSaveCardContent>): MsgSaveCardContent;
};
export declare const MsgSaveCardContentResponse: {
    encode(_: MsgSaveCardContentResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveCardContentResponse;
    fromJSON(_: any): MsgSaveCardContentResponse;
    toJSON(_: MsgSaveCardContentResponse): unknown;
    fromPartial(_: DeepPartial<MsgSaveCardContentResponse>): MsgSaveCardContentResponse;
};
export declare const MsgTransferCard: {
    encode(message: MsgTransferCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferCard;
    fromJSON(object: any): MsgTransferCard;
    toJSON(message: MsgTransferCard): unknown;
    fromPartial(object: DeepPartial<MsgTransferCard>): MsgTransferCard;
};
export declare const MsgTransferCardResponse: {
    encode(_: MsgTransferCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTransferCardResponse;
    fromJSON(_: any): MsgTransferCardResponse;
    toJSON(_: MsgTransferCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgTransferCardResponse>): MsgTransferCardResponse;
};
export declare const MsgDonateToCard: {
    encode(message: MsgDonateToCard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateToCard;
    fromJSON(object: any): MsgDonateToCard;
    toJSON(message: MsgDonateToCard): unknown;
    fromPartial(object: DeepPartial<MsgDonateToCard>): MsgDonateToCard;
};
export declare const MsgDonateToCardResponse: {
    encode(_: MsgDonateToCardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateToCardResponse;
    fromJSON(_: any): MsgDonateToCardResponse;
    toJSON(_: MsgDonateToCardResponse): unknown;
    fromPartial(_: DeepPartial<MsgDonateToCardResponse>): MsgDonateToCardResponse;
};
export declare const MsgAddArtwork: {
    encode(message: MsgAddArtwork, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddArtwork;
    fromJSON(object: any): MsgAddArtwork;
    toJSON(message: MsgAddArtwork): unknown;
    fromPartial(object: DeepPartial<MsgAddArtwork>): MsgAddArtwork;
};
export declare const MsgAddArtworkResponse: {
    encode(_: MsgAddArtworkResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddArtworkResponse;
    fromJSON(_: any): MsgAddArtworkResponse;
    toJSON(_: MsgAddArtworkResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddArtworkResponse>): MsgAddArtworkResponse;
};
export declare const MsgReportCopyrightViolation: {
    encode(message: MsgReportCopyrightViolation, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReportCopyrightViolation;
    fromJSON(object: any): MsgReportCopyrightViolation;
    toJSON(message: MsgReportCopyrightViolation): unknown;
    fromPartial(object: DeepPartial<MsgReportCopyrightViolation>): MsgReportCopyrightViolation;
};
export declare const MsgReportCopyrightViolationResponse: {
    encode(_: MsgReportCopyrightViolationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgReportCopyrightViolationResponse;
    fromJSON(_: any): MsgReportCopyrightViolationResponse;
    toJSON(_: MsgReportCopyrightViolationResponse): unknown;
    fromPartial(_: DeepPartial<MsgReportCopyrightViolationResponse>): MsgReportCopyrightViolationResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
    VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
    SaveCardContent(request: MsgSaveCardContent): Promise<MsgSaveCardContentResponse>;
    TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse>;
    DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse>;
    AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    ReportCopyrightViolation(request: MsgReportCopyrightViolation): Promise<MsgReportCopyrightViolationResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Createuser(request: MsgCreateuser): Promise<MsgCreateuserResponse>;
    BuyCardScheme(request: MsgBuyCardScheme): Promise<MsgBuyCardSchemeResponse>;
    VoteCard(request: MsgVoteCard): Promise<MsgVoteCardResponse>;
    SaveCardContent(request: MsgSaveCardContent): Promise<MsgSaveCardContentResponse>;
    TransferCard(request: MsgTransferCard): Promise<MsgTransferCardResponse>;
    DonateToCard(request: MsgDonateToCard): Promise<MsgDonateToCardResponse>;
    AddArtwork(request: MsgAddArtwork): Promise<MsgAddArtworkResponse>;
    ReportCopyrightViolation(request: MsgReportCopyrightViolation): Promise<MsgReportCopyrightViolationResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
