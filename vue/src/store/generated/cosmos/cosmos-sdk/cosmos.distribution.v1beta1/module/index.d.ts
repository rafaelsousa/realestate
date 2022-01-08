import { StdFee } from '@cosmjs/launchpad';
import { EncodeObject, OfflineSigner, Registry } from '@cosmjs/proto-signing';
import { Api } from './rest';
import { MsgFundCommunityPool, MsgSetWithdrawAddress, MsgWithdrawDelegatorReward, MsgWithdrawValidatorCommission } from './types/cosmos/distribution/v1beta1/tx';
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
    msgSetWithdrawAddress: (data: MsgSetWithdrawAddress) => EncodeObject;
    msgWithdrawDelegatorReward: (data: MsgWithdrawDelegatorReward) => EncodeObject;
    msgFundCommunityPool: (data: MsgFundCommunityPool) => EncodeObject;
    msgWithdrawValidatorCommission: (data: MsgWithdrawValidatorCommission) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
