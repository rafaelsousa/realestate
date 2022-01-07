// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateProperty } from "./types/realestate/tx";
import { MsgUpdateProperty } from "./types/realestate/tx";
import { MsgDeleteProperty } from "./types/realestate/tx";


const types = [
  ["/rafaelsousa.realestate.realestate.MsgCreateProperty", MsgCreateProperty],
  ["/rafaelsousa.realestate.realestate.MsgUpdateProperty", MsgUpdateProperty],
  ["/rafaelsousa.realestate.realestate.MsgDeleteProperty", MsgDeleteProperty],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateProperty: (data: MsgCreateProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.realestate.MsgCreateProperty", value: MsgCreateProperty.fromPartial( data ) }),
    msgUpdateProperty: (data: MsgUpdateProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.realestate.MsgUpdateProperty", value: MsgUpdateProperty.fromPartial( data ) }),
    msgDeleteProperty: (data: MsgDeleteProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.realestate.MsgDeleteProperty", value: MsgDeleteProperty.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
