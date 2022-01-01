/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rafaelsousa.realestate.realestate";

export interface Property {
  id: number;
  address: string;
  city: string;
  state: string;
  country: string;
  zipcode: string;
  latitude: string;
  longitude: string;
  creator: string;
}

const baseProperty: object = {
  id: 0,
  address: "",
  city: "",
  state: "",
  country: "",
  zipcode: "",
  latitude: "",
  longitude: "",
  creator: "",
};

export const Property = {
  encode(message: Property, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.city !== "") {
      writer.uint32(26).string(message.city);
    }
    if (message.state !== "") {
      writer.uint32(34).string(message.state);
    }
    if (message.country !== "") {
      writer.uint32(42).string(message.country);
    }
    if (message.zipcode !== "") {
      writer.uint32(50).string(message.zipcode);
    }
    if (message.latitude !== "") {
      writer.uint32(58).string(message.latitude);
    }
    if (message.longitude !== "") {
      writer.uint32(66).string(message.longitude);
    }
    if (message.creator !== "") {
      writer.uint32(74).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Property {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProperty } as Property;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.city = reader.string();
          break;
        case 4:
          message.state = reader.string();
          break;
        case 5:
          message.country = reader.string();
          break;
        case 6:
          message.zipcode = reader.string();
          break;
        case 7:
          message.latitude = reader.string();
          break;
        case 8:
          message.longitude = reader.string();
          break;
        case 9:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Property {
    const message = { ...baseProperty } as Property;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = String(object.city);
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
    }
    if (object.zipcode !== undefined && object.zipcode !== null) {
      message.zipcode = String(object.zipcode);
    } else {
      message.zipcode = "";
    }
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = String(object.latitude);
    } else {
      message.latitude = "";
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = String(object.longitude);
    } else {
      message.longitude = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Property): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.city !== undefined && (obj.city = message.city);
    message.state !== undefined && (obj.state = message.state);
    message.country !== undefined && (obj.country = message.country);
    message.zipcode !== undefined && (obj.zipcode = message.zipcode);
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Property>): Property {
    const message = { ...baseProperty } as Property;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.city !== undefined && object.city !== null) {
      message.city = object.city;
    } else {
      message.city = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
    }
    if (object.zipcode !== undefined && object.zipcode !== null) {
      message.zipcode = object.zipcode;
    } else {
      message.zipcode = "";
    }
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = object.latitude;
    } else {
      message.latitude = "";
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = object.longitude;
    } else {
      message.longitude = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

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
