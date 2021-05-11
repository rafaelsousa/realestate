// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateProperty } from "./types/property/tx";
import { MsgDeleteOwner } from "./types/property/tx";
import { MsgUpdateProperty } from "./types/property/tx";
import { MsgDeleteProperty } from "./types/property/tx";
import { MsgCreateOwner } from "./types/property/tx";
import { MsgUpdateOwner } from "./types/property/tx";


const types = [
  ["/rafaelsousa.realestate.property.MsgCreateProperty", MsgCreateProperty],
  ["/rafaelsousa.realestate.property.MsgDeleteOwner", MsgDeleteOwner],
  ["/rafaelsousa.realestate.property.MsgUpdateProperty", MsgUpdateProperty],
  ["/rafaelsousa.realestate.property.MsgDeleteProperty", MsgDeleteProperty],
  ["/rafaelsousa.realestate.property.MsgCreateOwner", MsgCreateOwner],
  ["/rafaelsousa.realestate.property.MsgUpdateOwner", MsgUpdateOwner],
  
];

const registry = new Registry(<any>types);

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
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgCreateProperty: (data: MsgCreateProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgCreateProperty", value: data }),
    msgDeleteOwner: (data: MsgDeleteOwner): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgDeleteOwner", value: data }),
    msgUpdateProperty: (data: MsgUpdateProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgUpdateProperty", value: data }),
    msgDeleteProperty: (data: MsgDeleteProperty): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgDeleteProperty", value: data }),
    msgCreateOwner: (data: MsgCreateOwner): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgCreateOwner", value: data }),
    msgUpdateOwner: (data: MsgUpdateOwner): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.property.MsgUpdateOwner", value: data }),
    
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
