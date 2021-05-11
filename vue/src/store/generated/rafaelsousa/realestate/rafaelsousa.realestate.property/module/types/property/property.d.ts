import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "rafaelsousa.realestate.property";
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
export declare const Property: {
    encode(message: Property, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Property;
    fromJSON(object: any): Property;
    toJSON(message: Property): unknown;
    fromPartial(object: DeepPartial<Property>): Property;
};
export declare const PropertyCollection: {
    encode(message: PropertyCollection, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PropertyCollection;
    fromJSON(object: any): PropertyCollection;
    toJSON(message: PropertyCollection): unknown;
    fromPartial(object: DeepPartial<PropertyCollection>): PropertyCollection;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
