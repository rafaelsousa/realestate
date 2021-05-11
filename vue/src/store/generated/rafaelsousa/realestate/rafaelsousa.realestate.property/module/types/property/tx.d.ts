import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "rafaelsousa.realestate.property";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateOwner {
    creator: string;
    address: string;
}
export interface MsgCreateOwnerResponse {
    id: number;
}
export interface MsgUpdateOwner {
    creator: string;
    id: number;
    address: string;
}
export interface MsgUpdateOwnerResponse {
}
export interface MsgDeleteOwner {
    creator: string;
    id: number;
}
export interface MsgDeleteOwnerResponse {
}
export interface MsgCreateProperty {
    creator: string;
    address: string;
    city: string;
    state: string;
    zip: string;
    country: string;
    latitude: string;
    longitude: string;
    ownerAddr: string;
}
export interface MsgCreatePropertyResponse {
    id: number;
}
export interface MsgUpdateProperty {
    creator: string;
    id: number;
    address: string;
    city: string;
    state: string;
    zip: string;
    country: string;
    latitude: string;
    longitude: string;
    ownerAddr: string;
}
export interface MsgUpdatePropertyResponse {
}
export interface MsgDeleteProperty {
    creator: string;
    id: number;
}
export interface MsgDeletePropertyResponse {
}
export declare const MsgCreateOwner: {
    encode(message: MsgCreateOwner, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateOwner;
    fromJSON(object: any): MsgCreateOwner;
    toJSON(message: MsgCreateOwner): unknown;
    fromPartial(object: DeepPartial<MsgCreateOwner>): MsgCreateOwner;
};
export declare const MsgCreateOwnerResponse: {
    encode(message: MsgCreateOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateOwnerResponse;
    fromJSON(object: any): MsgCreateOwnerResponse;
    toJSON(message: MsgCreateOwnerResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateOwnerResponse>): MsgCreateOwnerResponse;
};
export declare const MsgUpdateOwner: {
    encode(message: MsgUpdateOwner, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateOwner;
    fromJSON(object: any): MsgUpdateOwner;
    toJSON(message: MsgUpdateOwner): unknown;
    fromPartial(object: DeepPartial<MsgUpdateOwner>): MsgUpdateOwner;
};
export declare const MsgUpdateOwnerResponse: {
    encode(_: MsgUpdateOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateOwnerResponse;
    fromJSON(_: any): MsgUpdateOwnerResponse;
    toJSON(_: MsgUpdateOwnerResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateOwnerResponse>): MsgUpdateOwnerResponse;
};
export declare const MsgDeleteOwner: {
    encode(message: MsgDeleteOwner, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteOwner;
    fromJSON(object: any): MsgDeleteOwner;
    toJSON(message: MsgDeleteOwner): unknown;
    fromPartial(object: DeepPartial<MsgDeleteOwner>): MsgDeleteOwner;
};
export declare const MsgDeleteOwnerResponse: {
    encode(_: MsgDeleteOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteOwnerResponse;
    fromJSON(_: any): MsgDeleteOwnerResponse;
    toJSON(_: MsgDeleteOwnerResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteOwnerResponse>): MsgDeleteOwnerResponse;
};
export declare const MsgCreateProperty: {
    encode(message: MsgCreateProperty, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateProperty;
    fromJSON(object: any): MsgCreateProperty;
    toJSON(message: MsgCreateProperty): unknown;
    fromPartial(object: DeepPartial<MsgCreateProperty>): MsgCreateProperty;
};
export declare const MsgCreatePropertyResponse: {
    encode(message: MsgCreatePropertyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreatePropertyResponse;
    fromJSON(object: any): MsgCreatePropertyResponse;
    toJSON(message: MsgCreatePropertyResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreatePropertyResponse>): MsgCreatePropertyResponse;
};
export declare const MsgUpdateProperty: {
    encode(message: MsgUpdateProperty, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateProperty;
    fromJSON(object: any): MsgUpdateProperty;
    toJSON(message: MsgUpdateProperty): unknown;
    fromPartial(object: DeepPartial<MsgUpdateProperty>): MsgUpdateProperty;
};
export declare const MsgUpdatePropertyResponse: {
    encode(_: MsgUpdatePropertyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePropertyResponse;
    fromJSON(_: any): MsgUpdatePropertyResponse;
    toJSON(_: MsgUpdatePropertyResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdatePropertyResponse>): MsgUpdatePropertyResponse;
};
export declare const MsgDeleteProperty: {
    encode(message: MsgDeleteProperty, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteProperty;
    fromJSON(object: any): MsgDeleteProperty;
    toJSON(message: MsgDeleteProperty): unknown;
    fromPartial(object: DeepPartial<MsgDeleteProperty>): MsgDeleteProperty;
};
export declare const MsgDeletePropertyResponse: {
    encode(_: MsgDeletePropertyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePropertyResponse;
    fromJSON(_: any): MsgDeletePropertyResponse;
    toJSON(_: MsgDeletePropertyResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeletePropertyResponse>): MsgDeletePropertyResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateOwner(request: MsgCreateOwner): Promise<MsgCreateOwnerResponse>;
    UpdateOwner(request: MsgUpdateOwner): Promise<MsgUpdateOwnerResponse>;
    DeleteOwner(request: MsgDeleteOwner): Promise<MsgDeleteOwnerResponse>;
    CreateProperty(request: MsgCreateProperty): Promise<MsgCreatePropertyResponse>;
    UpdateProperty(request: MsgUpdateProperty): Promise<MsgUpdatePropertyResponse>;
    DeleteProperty(request: MsgDeleteProperty): Promise<MsgDeletePropertyResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateOwner(request: MsgCreateOwner): Promise<MsgCreateOwnerResponse>;
    UpdateOwner(request: MsgUpdateOwner): Promise<MsgUpdateOwnerResponse>;
    DeleteOwner(request: MsgDeleteOwner): Promise<MsgDeleteOwnerResponse>;
    CreateProperty(request: MsgCreateProperty): Promise<MsgCreatePropertyResponse>;
    UpdateProperty(request: MsgUpdateProperty): Promise<MsgUpdatePropertyResponse>;
    DeleteProperty(request: MsgDeleteProperty): Promise<MsgDeletePropertyResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
