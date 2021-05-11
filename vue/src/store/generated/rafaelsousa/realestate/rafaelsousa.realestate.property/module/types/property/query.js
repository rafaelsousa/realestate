/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Owner } from "../property/owner";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { Property } from "../property/property";
export const protobufPackage = "rafaelsousa.realestate.property";
const baseQueryGetOwnerRequest = { id: 0 };
export const QueryGetOwnerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetOwnerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetOwnerRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetOwnerRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseQueryGetOwnerResponse = {};
export const QueryGetOwnerResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Owner !== undefined) {
            Owner.encode(message.Owner, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetOwnerResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Owner = Owner.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetOwnerResponse };
        if (object.Owner !== undefined && object.Owner !== null) {
            message.Owner = Owner.fromJSON(object.Owner);
        }
        else {
            message.Owner = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Owner !== undefined &&
            (obj.Owner = message.Owner ? Owner.toJSON(message.Owner) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetOwnerResponse };
        if (object.Owner !== undefined && object.Owner !== null) {
            message.Owner = Owner.fromPartial(object.Owner);
        }
        else {
            message.Owner = undefined;
        }
        return message;
    },
};
const baseQueryAllOwnerRequest = {};
export const QueryAllOwnerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllOwnerRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllOwnerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllOwnerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllOwnerResponse = {};
export const QueryAllOwnerResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Owner) {
            Owner.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllOwnerResponse };
        message.Owner = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Owner.push(Owner.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryAllOwnerResponse };
        message.Owner = [];
        if (object.Owner !== undefined && object.Owner !== null) {
            for (const e of object.Owner) {
                message.Owner.push(Owner.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Owner) {
            obj.Owner = message.Owner.map((e) => (e ? Owner.toJSON(e) : undefined));
        }
        else {
            obj.Owner = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllOwnerResponse };
        message.Owner = [];
        if (object.Owner !== undefined && object.Owner !== null) {
            for (const e of object.Owner) {
                message.Owner.push(Owner.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryGetPropertyRequest = { id: 0 };
export const QueryGetPropertyRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetPropertyRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetPropertyRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetPropertyRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseQueryGetPropertyResponse = {};
export const QueryGetPropertyResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Property !== undefined) {
            Property.encode(message.Property, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetPropertyResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetPropertyResponse,
        };
        if (object.Property !== undefined && object.Property !== null) {
            message.Property = Property.fromJSON(object.Property);
        }
        else {
            message.Property = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Property !== undefined &&
            (obj.Property = message.Property
                ? Property.toJSON(message.Property)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetPropertyResponse,
        };
        if (object.Property !== undefined && object.Property !== null) {
            message.Property = Property.fromPartial(object.Property);
        }
        else {
            message.Property = undefined;
        }
        return message;
    },
};
const baseQueryAllPropertyRequest = {};
export const QueryAllPropertyRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllPropertyRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllPropertyRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllPropertyRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllPropertyResponse = {};
export const QueryAllPropertyResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Property) {
            Property.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllPropertyResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllPropertyResponse,
        };
        message.Property = [];
        if (object.Property !== undefined && object.Property !== null) {
            for (const e of object.Property) {
                message.Property.push(Property.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Property) {
            obj.Property = message.Property.map((e) => e ? Property.toJSON(e) : undefined);
        }
        else {
            obj.Property = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllPropertyResponse,
        };
        message.Property = [];
        if (object.Property !== undefined && object.Property !== null) {
            for (const e of object.Property) {
                message.Property.push(Property.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryListAllPropertiesRequest = { ownerAddr: "" };
export const QueryListAllPropertiesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.ownerAddr !== "") {
            writer.uint32(10).string(message.ownerAddr);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryListAllPropertiesRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.ownerAddr = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryListAllPropertiesRequest,
        };
        if (object.ownerAddr !== undefined && object.ownerAddr !== null) {
            message.ownerAddr = String(object.ownerAddr);
        }
        else {
            message.ownerAddr = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.ownerAddr !== undefined && (obj.ownerAddr = message.ownerAddr);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryListAllPropertiesRequest,
        };
        if (object.ownerAddr !== undefined && object.ownerAddr !== null) {
            message.ownerAddr = object.ownerAddr;
        }
        else {
            message.ownerAddr = "";
        }
        return message;
    },
};
const baseQueryListAllPropertiesResponse = {};
export const QueryListAllPropertiesResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Property) {
            Property.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryListAllPropertiesResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryListAllPropertiesResponse,
        };
        message.Property = [];
        if (object.Property !== undefined && object.Property !== null) {
            for (const e of object.Property) {
                message.Property.push(Property.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Property) {
            obj.Property = message.Property.map((e) => e ? Property.toJSON(e) : undefined);
        }
        else {
            obj.Property = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryListAllPropertiesResponse,
        };
        message.Property = [];
        if (object.Property !== undefined && object.Property !== null) {
            for (const e of object.Property) {
                message.Property.push(Property.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Owner(request) {
        const data = QueryGetOwnerRequest.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.property.Query", "Owner", data);
        return promise.then((data) => QueryGetOwnerResponse.decode(new Reader(data)));
    }
    OwnerAll(request) {
        const data = QueryAllOwnerRequest.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.property.Query", "OwnerAll", data);
        return promise.then((data) => QueryAllOwnerResponse.decode(new Reader(data)));
    }
    Property(request) {
        const data = QueryGetPropertyRequest.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.property.Query", "Property", data);
        return promise.then((data) => QueryGetPropertyResponse.decode(new Reader(data)));
    }
    PropertyAll(request) {
        const data = QueryAllPropertyRequest.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.property.Query", "PropertyAll", data);
        return promise.then((data) => QueryAllPropertyResponse.decode(new Reader(data)));
    }
    AllProperties(request) {
        const data = QueryListAllPropertiesRequest.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.property.Query", "AllProperties", data);
        return promise.then((data) => QueryListAllPropertiesResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
