/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../blueprints/params";
import { House } from "../blueprints/house";
export const protobufPackage = "rafaelsousa.realestate.blueprints";
const baseGenesisState = { houseCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.houseList) {
            House.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.houseCount !== 0) {
            writer.uint32(24).uint64(message.houseCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.houseList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.houseList.push(House.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.houseCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.houseList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.houseList !== undefined && object.houseList !== null) {
            for (const e of object.houseList) {
                message.houseList.push(House.fromJSON(e));
            }
        }
        if (object.houseCount !== undefined && object.houseCount !== null) {
            message.houseCount = Number(object.houseCount);
        }
        else {
            message.houseCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.houseList) {
            obj.houseList = message.houseList.map((e) => e ? House.toJSON(e) : undefined);
        }
        else {
            obj.houseList = [];
        }
        message.houseCount !== undefined && (obj.houseCount = message.houseCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.houseList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.houseList !== undefined && object.houseList !== null) {
            for (const e of object.houseList) {
                message.houseList.push(House.fromPartial(e));
            }
        }
        if (object.houseCount !== undefined && object.houseCount !== null) {
            message.houseCount = object.houseCount;
        }
        else {
            message.houseCount = 0;
        }
        return message;
    },
};
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
