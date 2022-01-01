import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "rafaelsousa.realestate.realestate";
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
export declare const Property: {
    encode(message: Property, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Property;
    fromJSON(object: any): Property;
    toJSON(message: Property): unknown;
    fromPartial(object: DeepPartial<Property>): Property;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
