import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../realestate/params";
import { Property } from "../realestate/property";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "rafaelsousa.realestate.realestate";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
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
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
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
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Property by id. */
    Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
    /** Queries a list of Property items. */
    PropertyAll(request: QueryAllPropertyRequest): Promise<QueryAllPropertyResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
    PropertyAll(request: QueryAllPropertyRequest): Promise<QueryAllPropertyResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
