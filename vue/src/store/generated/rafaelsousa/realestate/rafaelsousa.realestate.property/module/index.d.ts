import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateProperty } from "./types/property/tx";
import { MsgDeleteProperty } from "./types/property/tx";
import { MsgCreateProperty } from "./types/property/tx";
import { MsgUpdateOwner } from "./types/property/tx";
import { MsgCreateOwner } from "./types/property/tx";
import { MsgDeleteOwner } from "./types/property/tx";
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgUpdateProperty: (data: MsgUpdateProperty) => EncodeObject;
    msgDeleteProperty: (data: MsgDeleteProperty) => EncodeObject;
    msgCreateProperty: (data: MsgCreateProperty) => EncodeObject;
    msgUpdateOwner: (data: MsgUpdateOwner) => EncodeObject;
    msgCreateOwner: (data: MsgCreateOwner) => EncodeObject;
    msgDeleteOwner: (data: MsgDeleteOwner) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
