import { Reader, Writer } from "protobufjs/minimal";
import { Property } from "../property/property";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "rafaelsousa.realestate.property";
/** this line is used by starport scaffolding # 3 */
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
    /** this line is used by starport scaffolding # 2 */
    Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
    PropertyAll(request: QueryAllPropertyRequest): Promise<QueryAllPropertyResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
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
