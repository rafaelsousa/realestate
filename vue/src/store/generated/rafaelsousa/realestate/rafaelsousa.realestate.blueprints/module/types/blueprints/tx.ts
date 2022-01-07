/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "rafaelsousa.realestate.blueprints";

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

export interface MsgUpdateHouseResponse {}

export interface MsgDeleteHouse {
  creator: string;
  id: number;
}

export interface MsgDeleteHouseResponse {}

const baseMsgCreateHouse: object = { creator: "", description: "", image: "" };

export const MsgCreateHouse = {
  encode(message: MsgCreateHouse, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.image !== "") {
      writer.uint32(26).string(message.image);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateHouse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateHouse } as MsgCreateHouse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.image = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateHouse {
    const message = { ...baseMsgCreateHouse } as MsgCreateHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    return message;
  },

  toJSON(message: MsgCreateHouse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.description !== undefined &&
      (obj.description = message.description);
    message.image !== undefined && (obj.image = message.image);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateHouse>): MsgCreateHouse {
    const message = { ...baseMsgCreateHouse } as MsgCreateHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    return message;
  },
};

const baseMsgCreateHouseResponse: object = { id: 0 };

export const MsgCreateHouseResponse = {
  encode(
    message: MsgCreateHouseResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateHouseResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateHouseResponse } as MsgCreateHouseResponse;
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

  fromJSON(object: any): MsgCreateHouseResponse {
    const message = { ...baseMsgCreateHouseResponse } as MsgCreateHouseResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateHouseResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateHouseResponse>
  ): MsgCreateHouseResponse {
    const message = { ...baseMsgCreateHouseResponse } as MsgCreateHouseResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateHouse: object = {
  creator: "",
  id: 0,
  description: "",
  image: "",
};

export const MsgUpdateHouse = {
  encode(message: MsgUpdateHouse, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.image !== "") {
      writer.uint32(34).string(message.image);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateHouse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateHouse } as MsgUpdateHouse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.image = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateHouse {
    const message = { ...baseMsgUpdateHouse } as MsgUpdateHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateHouse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.description !== undefined &&
      (obj.description = message.description);
    message.image !== undefined && (obj.image = message.image);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateHouse>): MsgUpdateHouse {
    const message = { ...baseMsgUpdateHouse } as MsgUpdateHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    return message;
  },
};

const baseMsgUpdateHouseResponse: object = {};

export const MsgUpdateHouseResponse = {
  encode(_: MsgUpdateHouseResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateHouseResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateHouseResponse } as MsgUpdateHouseResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateHouseResponse {
    const message = { ...baseMsgUpdateHouseResponse } as MsgUpdateHouseResponse;
    return message;
  },

  toJSON(_: MsgUpdateHouseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateHouseResponse>): MsgUpdateHouseResponse {
    const message = { ...baseMsgUpdateHouseResponse } as MsgUpdateHouseResponse;
    return message;
  },
};

const baseMsgDeleteHouse: object = { creator: "", id: 0 };

export const MsgDeleteHouse = {
  encode(message: MsgDeleteHouse, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteHouse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteHouse } as MsgDeleteHouse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteHouse {
    const message = { ...baseMsgDeleteHouse } as MsgDeleteHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteHouse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteHouse>): MsgDeleteHouse {
    const message = { ...baseMsgDeleteHouse } as MsgDeleteHouse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteHouseResponse: object = {};

export const MsgDeleteHouseResponse = {
  encode(_: MsgDeleteHouseResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteHouseResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteHouseResponse } as MsgDeleteHouseResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteHouseResponse {
    const message = { ...baseMsgDeleteHouseResponse } as MsgDeleteHouseResponse;
    return message;
  },

  toJSON(_: MsgDeleteHouseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteHouseResponse>): MsgDeleteHouseResponse {
    const message = { ...baseMsgDeleteHouseResponse } as MsgDeleteHouseResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateHouse(request: MsgCreateHouse): Promise<MsgCreateHouseResponse>;
  UpdateHouse(request: MsgUpdateHouse): Promise<MsgUpdateHouseResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteHouse(request: MsgDeleteHouse): Promise<MsgDeleteHouseResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateHouse(request: MsgCreateHouse): Promise<MsgCreateHouseResponse> {
    const data = MsgCreateHouse.encode(request).finish();
    const promise = this.rpc.request(
      "rafaelsousa.realestate.blueprints.Msg",
      "CreateHouse",
      data
    );
    return promise.then((data) =>
      MsgCreateHouseResponse.decode(new Reader(data))
    );
  }

  UpdateHouse(request: MsgUpdateHouse): Promise<MsgUpdateHouseResponse> {
    const data = MsgUpdateHouse.encode(request).finish();
    const promise = this.rpc.request(
      "rafaelsousa.realestate.blueprints.Msg",
      "UpdateHouse",
      data
    );
    return promise.then((data) =>
      MsgUpdateHouseResponse.decode(new Reader(data))
    );
  }

  DeleteHouse(request: MsgDeleteHouse): Promise<MsgDeleteHouseResponse> {
    const data = MsgDeleteHouse.encode(request).finish();
    const promise = this.rpc.request(
      "rafaelsousa.realestate.blueprints.Msg",
      "DeleteHouse",
      data
    );
    return promise.then((data) =>
      MsgDeleteHouseResponse.decode(new Reader(data))
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
