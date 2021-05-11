/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rafaelsousa.realestate.property";

export interface Property {
  creator: string;
  id: number;
  address: string;
  city: string;
  state: string;
  zip: string;
  country: string;
  latitude: string;
  longitude: string;
  owneraddr: string;
}

export interface PropertyCollection {
  properties: Property[];
}

const baseProperty: object = {
  creator: "",
  id: 0,
  address: "",
  city: "",
  state: "",
  zip: "",
  country: "",
  latitude: "",
  longitude: "",
  owneraddr: "",
};

export const Property = {
  encode(message: Property, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.city !== "") {
      writer.uint32(34).string(message.city);
    }
    if (message.state !== "") {
      writer.uint32(42).string(message.state);
    }
    if (message.zip !== "") {
      writer.uint32(50).string(message.zip);
    }
    if (message.country !== "") {
      writer.uint32(58).string(message.country);
    }
    if (message.latitude !== "") {
      writer.uint32(66).string(message.latitude);
    }
    if (message.longitude !== "") {
      writer.uint32(74).string(message.longitude);
    }
    if (message.owneraddr !== "") {
      writer.uint32(82).string(message.owneraddr);
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
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.city = reader.string();
          break;
        case 5:
          message.state = reader.string();
          break;
        case 6:
          message.zip = reader.string();
          break;
        case 7:
          message.country = reader.string();
          break;
        case 8:
          message.latitude = reader.string();
          break;
        case 9:
          message.longitude = reader.string();
          break;
        case 10:
          message.owneraddr = reader.string();
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
    if (object.zip !== undefined && object.zip !== null) {
      message.zip = String(object.zip);
    } else {
      message.zip = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = String(object.country);
    } else {
      message.country = "";
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
    if (object.owneraddr !== undefined && object.owneraddr !== null) {
      message.owneraddr = String(object.owneraddr);
    } else {
      message.owneraddr = "";
    }
    return message;
  },

  toJSON(message: Property): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.city !== undefined && (obj.city = message.city);
    message.state !== undefined && (obj.state = message.state);
    message.zip !== undefined && (obj.zip = message.zip);
    message.country !== undefined && (obj.country = message.country);
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    message.owneraddr !== undefined && (obj.owneraddr = message.owneraddr);
    return obj;
  },

  fromPartial(object: DeepPartial<Property>): Property {
    const message = { ...baseProperty } as Property;
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
    if (object.zip !== undefined && object.zip !== null) {
      message.zip = object.zip;
    } else {
      message.zip = "";
    }
    if (object.country !== undefined && object.country !== null) {
      message.country = object.country;
    } else {
      message.country = "";
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
    if (object.owneraddr !== undefined && object.owneraddr !== null) {
      message.owneraddr = object.owneraddr;
    } else {
      message.owneraddr = "";
    }
    return message;
  },
};

const basePropertyCollection: object = {};

export const PropertyCollection = {
  encode(
    message: PropertyCollection,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.properties) {
      Property.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PropertyCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePropertyCollection } as PropertyCollection;
    message.properties = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.properties.push(Property.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PropertyCollection {
    const message = { ...basePropertyCollection } as PropertyCollection;
    message.properties = [];
    if (object.properties !== undefined && object.properties !== null) {
      for (const e of object.properties) {
        message.properties.push(Property.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: PropertyCollection): unknown {
    const obj: any = {};
    if (message.properties) {
      obj.properties = message.properties.map((e) =>
        e ? Property.toJSON(e) : undefined
      );
    } else {
      obj.properties = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<PropertyCollection>): PropertyCollection {
    const message = { ...basePropertyCollection } as PropertyCollection;
    message.properties = [];
    if (object.properties !== undefined && object.properties !== null) {
      for (const e of object.properties) {
        message.properties.push(Property.fromPartial(e));
      }
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
