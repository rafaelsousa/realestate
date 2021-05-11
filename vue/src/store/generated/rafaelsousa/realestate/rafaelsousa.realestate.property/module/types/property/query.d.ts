import { Reader, Writer } from "protobufjs/minimal";
import { Owner } from "../property/owner";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Property } from "../property/property";
export declare const protobufPackage = "rafaelsousa.realestate.property";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetOwnerRequest {
    id: number;
}
export interface QueryGetOwnerResponse {
    Owner: Owner | undefined;
}
export interface QueryAllOwnerRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllOwnerResponse {
    Owner: Owner[];
    pagination: PageResponse | undefined;
}
export interface QueryGetPropertyRequest {
    id: number;
}
export interface QueryGetPropertyResponse {
    Property: Property | undefined;
}
export interface QueryAllPropertyRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPropertyResponse {
    Property: Property[];
    pagination: PageResponse | undefined;
}
export interface QueryListAllPropertiesRequest {
    ownerAddr: string;
}
export interface QueryListAllPropertiesResponse {
    Property: Property[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetOwnerRequest: {
    encode(message: QueryGetOwnerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerRequest;
    fromJSON(object: any): QueryGetOwnerRequest;
    toJSON(message: QueryGetOwnerRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetOwnerRequest>): QueryGetOwnerRequest;
};
export declare const QueryGetOwnerResponse: {
    encode(message: QueryGetOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerResponse;
    fromJSON(object: any): QueryGetOwnerResponse;
    toJSON(message: QueryGetOwnerResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetOwnerResponse>): QueryGetOwnerResponse;
};
export declare const QueryAllOwnerRequest: {
    encode(message: QueryAllOwnerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOwnerRequest;
    fromJSON(object: any): QueryAllOwnerRequest;
    toJSON(message: QueryAllOwnerRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllOwnerRequest>): QueryAllOwnerRequest;
};
export declare const QueryAllOwnerResponse: {
    encode(message: QueryAllOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllOwnerResponse;
    fromJSON(object: any): QueryAllOwnerResponse;
    toJSON(message: QueryAllOwnerResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllOwnerResponse>): QueryAllOwnerResponse;
};
export declare const QueryGetPropertyRequest: {
    encode(message: QueryGetPropertyRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPropertyRequest;
    fromJSON(object: any): QueryGetPropertyRequest;
    toJSON(message: QueryGetPropertyRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPropertyRequest>): QueryGetPropertyRequest;
};
export declare const QueryGetPropertyResponse: {
    encode(message: QueryGetPropertyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPropertyResponse;
    fromJSON(object: any): QueryGetPropertyResponse;
    toJSON(message: QueryGetPropertyResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPropertyResponse>): QueryGetPropertyResponse;
};
export declare const QueryAllPropertyRequest: {
    encode(message: QueryAllPropertyRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPropertyRequest;
    fromJSON(object: any): QueryAllPropertyRequest;
    toJSON(message: QueryAllPropertyRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPropertyRequest>): QueryAllPropertyRequest;
};
export declare const QueryAllPropertyResponse: {
    encode(message: QueryAllPropertyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPropertyResponse;
    fromJSON(object: any): QueryAllPropertyResponse;
    toJSON(message: QueryAllPropertyResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPropertyResponse>): QueryAllPropertyResponse;
};
export declare const QueryListAllPropertiesRequest: {
    encode(message: QueryListAllPropertiesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryListAllPropertiesRequest;
    fromJSON(object: any): QueryListAllPropertiesRequest;
    toJSON(message: QueryListAllPropertiesRequest): unknown;
    fromPartial(object: DeepPartial<QueryListAllPropertiesRequest>): QueryListAllPropertiesRequest;
};
export declare const QueryListAllPropertiesResponse: {
    encode(message: QueryListAllPropertiesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryListAllPropertiesResponse;
    fromJSON(object: any): QueryListAllPropertiesResponse;
    toJSON(message: QueryListAllPropertiesResponse): unknown;
    fromPartial(object: DeepPartial<QueryListAllPropertiesResponse>): QueryListAllPropertiesResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** this line is used by starport scaffolding # 2 */
    Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse>;
    OwnerAll(request: QueryAllOwnerRequest): Promise<QueryAllOwnerResponse>;
    Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
    PropertyAll(request: QueryAllPropertyRequest): Promise<QueryAllPropertyResponse>;
    /** Returns all properties from a specific owner */
    AllProperties(request: QueryListAllPropertiesRequest): Promise<QueryListAllPropertiesResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse>;
    OwnerAll(request: QueryAllOwnerRequest): Promise<QueryAllOwnerResponse>;
    Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
    PropertyAll(request: QueryAllPropertyRequest): Promise<QueryAllPropertyResponse>;
    AllProperties(request: QueryListAllPropertiesRequest): Promise<QueryListAllPropertiesResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
