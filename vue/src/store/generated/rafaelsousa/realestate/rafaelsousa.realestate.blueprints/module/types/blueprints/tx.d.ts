import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "rafaelsousa.realestate.blueprints";
export interface MsgCreateHouse {
    creator: string;
    description: string;
    image: string;
}
export interface MsgCreateHouseResponse {
    id: number;
}
export interface MsgUpdateHouse {
    creator: string;
    id: number;
    description: string;
    image: string;
}
export interface MsgUpdateHouseResponse {
}
export interface MsgDeleteHouse {
    creator: string;
    id: number;
}
export interface MsgDeleteHouseResponse {
}
export declare const MsgCreateHouse: {
    encode(message: MsgCreateHouse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHouse;
    fromJSON(object: any): MsgCreateHouse;
    toJSON(message: MsgCreateHouse): unknown;
    fromPartial(object: DeepPartial<MsgCreateHouse>): MsgCreateHouse;
};
export declare const MsgCreateHouseResponse: {
    encode(message: MsgCreateHouseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHouseResponse;
    fromJSON(object: any): MsgCreateHouseResponse;
    toJSON(message: MsgCreateHouseResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateHouseResponse>): MsgCreateHouseResponse;
};
export declare const MsgUpdateHouse: {
    encode(message: MsgUpdateHouse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateHouse;
    fromJSON(object: any): MsgUpdateHouse;
    toJSON(message: MsgUpdateHouse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateHouse>): MsgUpdateHouse;
};
export declare const MsgUpdateHouseResponse: {
    encode(_: MsgUpdateHouseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateHouseResponse;
    fromJSON(_: any): MsgUpdateHouseResponse;
    toJSON(_: MsgUpdateHouseResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateHouseResponse>): MsgUpdateHouseResponse;
};
export declare const MsgDeleteHouse: {
    encode(message: MsgDeleteHouse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteHouse;
    fromJSON(object: any): MsgDeleteHouse;
    toJSON(message: MsgDeleteHouse): unknown;
    fromPartial(object: DeepPartial<MsgDeleteHouse>): MsgDeleteHouse;
};
export declare const MsgDeleteHouseResponse: {
    encode(_: MsgDeleteHouseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteHouseResponse;
    fromJSON(_: any): MsgDeleteHouseResponse;
    toJSON(_: MsgDeleteHouseResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteHouseResponse>): MsgDeleteHouseResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateHouse(request: MsgCreateHouse): Promise<MsgCreateHouseResponse>;
    UpdateHouse(request: MsgUpdateHouse): Promise<MsgUpdateHouseResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteHouse(request: MsgDeleteHouse): Promise<MsgDeleteHouseResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateHouse(request: MsgCreateHouse): Promise<MsgCreateHouseResponse>;
    UpdateHouse(request: MsgUpdateHouse): Promise<MsgUpdateHouseResponse>;
    DeleteHouse(request: MsgDeleteHouse): Promise<MsgDeleteHouseResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
