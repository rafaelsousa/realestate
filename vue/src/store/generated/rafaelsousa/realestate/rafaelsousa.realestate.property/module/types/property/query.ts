/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Property } from "../property/property";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "rafaelsousa.realestate.property";

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

const baseQueryGetPropertyRequest: object = { id: 0 };

export const QueryGetPropertyRequest = {
  encode(
    message: QueryGetPropertyRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPropertyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPropertyRequest,
    } as QueryGetPropertyRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPropertyRequest {
    const message = {
      ...baseQueryGetPropertyRequest,
    } as QueryGetPropertyRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetPropertyRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPropertyRequest>
  ): QueryGetPropertyRequest {
    const message = {
      ...baseQueryGetPropertyRequest,
    } as QueryGetPropertyRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetPropertyResponse: object = {};

export const QueryGetPropertyResponse = {
  encode(
    message: QueryGetPropertyResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Property !== undefined) {
      Property.encode(message.Property, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPropertyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPropertyResponse,
    } as QueryGetPropertyResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Property = Property.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPropertyResponse {
    const message = {
      ...baseQueryGetPropertyResponse,
    } as QueryGetPropertyResponse;
    if (object.Property !== undefined && object.Property !== null) {
      message.Property = Property.fromJSON(object.Property);
    } else {
      message.Property = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPropertyResponse): unknown {
    const obj: any = {};
    message.Property !== undefined &&
      (obj.Property = message.Property
        ? Property.toJSON(message.Property)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPropertyResponse>
  ): QueryGetPropertyResponse {
    const message = {
      ...baseQueryGetPropertyResponse,
    } as QueryGetPropertyResponse;
    if (object.Property !== undefined && object.Property !== null) {
      message.Property = Property.fromPartial(object.Property);
    } else {
      message.Property = undefined;
    }
    return message;
  },
};

const baseQueryAllPropertyRequest: object = {};

export const QueryAllPropertyRequest = {
  encode(
    message: QueryAllPropertyRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPropertyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPropertyRequest,
    } as QueryAllPropertyRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPropertyRequest {
    const message = {
      ...baseQueryAllPropertyRequest,
    } as QueryAllPropertyRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPropertyRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPropertyRequest>
  ): QueryAllPropertyRequest {
    const message = {
      ...baseQueryAllPropertyRequest,
    } as QueryAllPropertyRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPropertyResponse: object = {};

export const QueryAllPropertyResponse = {
  encode(
    message: QueryAllPropertyResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Property) {
      Property.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPropertyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPropertyResponse,
    } as QueryAllPropertyResponse;
    message.Property = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Property.push(Property.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPropertyResponse {
    const message = {
      ...baseQueryAllPropertyResponse,
    } as QueryAllPropertyResponse;
    message.Property = [];
    if (object.Property !== undefined && object.Property !== null) {
      for (const e of object.Property) {
        message.Property.push(Property.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPropertyResponse): unknown {
    const obj: any = {};
    if (message.Property) {
      obj.Property = message.Property.map((e) =>
        e ? Property.toJSON(e) : undefined
      );
    } else {
      obj.Property = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPropertyResponse>
  ): QueryAllPropertyResponse {
    const message = {
      ...baseQueryAllPropertyResponse,
    } as QueryAllPropertyResponse;
    message.Property = [];
    if (object.Property !== undefined && object.Property !== null) {
      for (const e of object.Property) {
        message.Property.push(Property.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** this line is used by starport scaffolding # 2 */
  Property(request: QueryGetPropertyRequest): Promise<QueryGetPropertyResponse>;
  PropertyAll(
    request: QueryAllPropertyRequest
  ): Promise<QueryAllPropertyResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Property(
    request: QueryGetPropertyRequest
  ): Promise<QueryGetPropertyResponse> {
    const data = QueryGetPropertyRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rafaelsousa.realestate.property.Query",
      "Property",
      data
    );
    return promise.then((data) =>
      QueryGetPropertyResponse.decode(new Reader(data))
    );
  }

  PropertyAll(
    request: QueryAllPropertyRequest
  ): Promise<QueryAllPropertyResponse> {
    const data = QueryAllPropertyRequest.encode(request).finish();
    const promise = this.rpc.request(
      "rafaelsousa.realestate.property.Query",
      "PropertyAll",
      data
    );
    return promise.then((data) =>
      QueryAllPropertyResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
