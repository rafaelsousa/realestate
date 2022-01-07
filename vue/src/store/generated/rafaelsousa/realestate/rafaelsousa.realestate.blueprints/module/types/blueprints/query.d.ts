import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../blueprints/params";
import { House } from "../blueprints/house";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "rafaelsousa.realestate.blueprints";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetHouseRequest {
    id: number;
}
export interface QueryGetHouseResponse {
    House: House | undefined;
}
export interface QueryAllHouseRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllHouseResponse {
    House: House[];
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
export declare const QueryGetHouseRequest: {
    encode(message: QueryGetHouseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetHouseRequest;
    fromJSON(object: any): QueryGetHouseRequest;
    toJSON(message: QueryGetHouseRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetHouseRequest>): QueryGetHouseRequest;
};
export declare const QueryGetHouseResponse: {
    encode(message: QueryGetHouseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetHouseResponse;
    fromJSON(object: any): QueryGetHouseResponse;
    toJSON(message: QueryGetHouseResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetHouseResponse>): QueryGetHouseResponse;
};
export declare const QueryAllHouseRequest: {
    encode(message: QueryAllHouseRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllHouseRequest;
    fromJSON(object: any): QueryAllHouseRequest;
    toJSON(message: QueryAllHouseRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllHouseRequest>): QueryAllHouseRequest;
};
export declare const QueryAllHouseResponse: {
    encode(message: QueryAllHouseResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllHouseResponse;
    fromJSON(object: any): QueryAllHouseResponse;
    toJSON(message: QueryAllHouseResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllHouseResponse>): QueryAllHouseResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a House by id. */
    House(request: QueryGetHouseRequest): Promise<QueryGetHouseResponse>;
    /** Queries a list of House items. */
    HouseAll(request: QueryAllHouseRequest): Promise<QueryAllHouseResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    House(request: QueryGetHouseRequest): Promise<QueryGetHouseResponse>;
    HouseAll(request: QueryAllHouseRequest): Promise<QueryAllHouseResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
