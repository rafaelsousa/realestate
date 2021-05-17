/* eslint-disable */
import { Property } from "../property/property";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "rafaelsousa.realestate.property";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.propertyList) {
            Property.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.propertyList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.propertyList.push(Property.decode(reader, reader.uint32()));
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
        message.propertyList = [];
        if (object.propertyList !== undefined && object.propertyList !== null) {
            for (const e of object.propertyList) {
                message.propertyList.push(Property.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.propertyList) {
            obj.propertyList = message.propertyList.map((e) => e ? Property.toJSON(e) : undefined);
        }
        else {
            obj.propertyList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.propertyList = [];
        if (object.propertyList !== undefined && object.propertyList !== null) {
            for (const e of object.propertyList) {
                message.propertyList.push(Property.fromPartial(e));
            }
        }
        return message;
    },
};
