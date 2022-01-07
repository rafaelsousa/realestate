import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "rafaelsousa.realestate.blueprints";
export interface House {
    id: number;
    description: string;
    image: string;
    creator: string;
}
export declare const House: {
    encode(message: House, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): House;
    fromJSON(object: any): House;
    toJSON(message: House): unknown;
    fromPartial(object: DeepPartial<House>): House;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
