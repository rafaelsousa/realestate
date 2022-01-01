/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "rafaelsousa.realestate.realestate";
const baseMsgCreateProperty = {
    creator: "",
    address: "",
    city: "",
    state: "",
    country: "",
    zipcode: "",
    latitude: "",
    longitude: "",
};
export const MsgCreateProperty = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateProperty };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.city !== undefined && object.city !== null) {
            message.city = String(object.city);
        }
        else {
            message.city = "";
        }
        if (object.state !== undefined && object.state !== null) {
            message.state = String(object.state);
        }
        else {
            message.state = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = String(object.country);
        }
        else {
            message.country = "";
        }
        if (object.zipcode !== undefined && object.zipcode !== null) {
            message.zipcode = String(object.zipcode);
        }
        else {
            message.zipcode = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = String(object.latitude);
        }
        else {
            message.latitude = "";
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = String(object.longitude);
        }
        else {
            message.longitude = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.address !== undefined && (obj.address = message.address);
        message.city !== undefined && (obj.city = message.city);
        message.state !== undefined && (obj.state = message.state);
        message.country !== undefined && (obj.country = message.country);
        message.zipcode !== undefined && (obj.zipcode = message.zipcode);
        message.latitude !== undefined && (obj.latitude = message.latitude);
        message.longitude !== undefined && (obj.longitude = message.longitude);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.city !== undefined && object.city !== null) {
            message.city = object.city;
        }
        else {
            message.city = "";
        }
        if (object.state !== undefined && object.state !== null) {
            message.state = object.state;
        }
        else {
            message.state = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = object.country;
        }
        else {
            message.country = "";
        }
        if (object.zipcode !== undefined && object.zipcode !== null) {
            message.zipcode = object.zipcode;
        }
        else {
            message.zipcode = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = object.latitude;
        }
        else {
            message.latitude = "";
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = object.longitude;
        }
        else {
            message.longitude = "";
        }
        return message;
    },
};
const baseMsgCreatePropertyResponse = { id: 0 };
export const MsgCreatePropertyResponse = {
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
            ...baseMsgCreatePropertyResponse,
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
            ...baseMsgCreatePropertyResponse,
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
            ...baseMsgCreatePropertyResponse,
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
const baseMsgUpdateProperty = {
    creator: "",
    id: 0,
    address: "",
    city: "",
    state: "",
    country: "",
    zipcode: "",
    latitude: "",
    longitude: "",
};
export const MsgUpdateProperty = {
    encode(message, writer = Writer.create()) {
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
        if (message.country !== "") {
            writer.uint32(50).string(message.country);
        }
        if (message.zipcode !== "") {
            writer.uint32(58).string(message.zipcode);
        }
        if (message.latitude !== "") {
            writer.uint32(66).string(message.latitude);
        }
        if (message.longitude !== "") {
            writer.uint32(74).string(message.longitude);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateProperty };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
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
                    message.country = reader.string();
                    break;
                case 7:
                    message.zipcode = reader.string();
                    break;
                case 8:
                    message.latitude = reader.string();
                    break;
                case 9:
                    message.longitude = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.city !== undefined && object.city !== null) {
            message.city = String(object.city);
        }
        else {
            message.city = "";
        }
        if (object.state !== undefined && object.state !== null) {
            message.state = String(object.state);
        }
        else {
            message.state = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = String(object.country);
        }
        else {
            message.country = "";
        }
        if (object.zipcode !== undefined && object.zipcode !== null) {
            message.zipcode = String(object.zipcode);
        }
        else {
            message.zipcode = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = String(object.latitude);
        }
        else {
            message.latitude = "";
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = String(object.longitude);
        }
        else {
            message.longitude = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        message.address !== undefined && (obj.address = message.address);
        message.city !== undefined && (obj.city = message.city);
        message.state !== undefined && (obj.state = message.state);
        message.country !== undefined && (obj.country = message.country);
        message.zipcode !== undefined && (obj.zipcode = message.zipcode);
        message.latitude !== undefined && (obj.latitude = message.latitude);
        message.longitude !== undefined && (obj.longitude = message.longitude);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.city !== undefined && object.city !== null) {
            message.city = object.city;
        }
        else {
            message.city = "";
        }
        if (object.state !== undefined && object.state !== null) {
            message.state = object.state;
        }
        else {
            message.state = "";
        }
        if (object.country !== undefined && object.country !== null) {
            message.country = object.country;
        }
        else {
            message.country = "";
        }
        if (object.zipcode !== undefined && object.zipcode !== null) {
            message.zipcode = object.zipcode;
        }
        else {
            message.zipcode = "";
        }
        if (object.latitude !== undefined && object.latitude !== null) {
            message.latitude = object.latitude;
        }
        else {
            message.latitude = "";
        }
        if (object.longitude !== undefined && object.longitude !== null) {
            message.longitude = object.longitude;
        }
        else {
            message.longitude = "";
        }
        return message;
    },
};
const baseMsgUpdatePropertyResponse = {};
export const MsgUpdatePropertyResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdatePropertyResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdatePropertyResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdatePropertyResponse,
        };
        return message;
    },
};
const baseMsgDeleteProperty = { creator: "", id: 0 };
export const MsgDeleteProperty = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.id !== 0) {
            writer.uint32(16).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteProperty };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgDeleteProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteProperty };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgDeletePropertyResponse = {};
export const MsgDeletePropertyResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeletePropertyResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgDeletePropertyResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgDeletePropertyResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateProperty(request) {
        const data = MsgCreateProperty.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.realestate.Msg", "CreateProperty", data);
        return promise.then((data) => MsgCreatePropertyResponse.decode(new Reader(data)));
    }
    UpdateProperty(request) {
        const data = MsgUpdateProperty.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.realestate.Msg", "UpdateProperty", data);
        return promise.then((data) => MsgUpdatePropertyResponse.decode(new Reader(data)));
    }
    DeleteProperty(request) {
        const data = MsgDeleteProperty.encode(request).finish();
        const promise = this.rpc.request("rafaelsousa.realestate.realestate.Msg", "DeleteProperty", data);
        return promise.then((data) => MsgDeletePropertyResponse.decode(new Reader(data)));
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
