import { StdFee } from '@cosmjs/launchpad';
import { EncodeObject, OfflineSigner, Registry } from '@cosmjs/proto-signing';
import { Api } from './rest';
import { MsgBeginRedelegate, MsgCreateValidator, MsgDelegate, MsgEditValidator, MsgUndelegate } from './types/cosmos/staking/v1beta1/tx';
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgEditValidator: (data: MsgEditValidator) => EncodeObject;
    msgCreateValidator: (data: MsgCreateValidator) => EncodeObject;
    msgDelegate: (data: MsgDelegate) => EncodeObject;
    msgBeginRedelegate: (data: MsgBeginRedelegate) => EncodeObject;
    msgUndelegate: (data: MsgUndelegate) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
