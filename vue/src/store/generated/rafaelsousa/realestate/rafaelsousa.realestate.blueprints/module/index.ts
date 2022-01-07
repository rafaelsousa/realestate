// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteHouse } from "./types/blueprints/tx";
import { MsgCreateHouse } from "./types/blueprints/tx";
import { MsgUpdateHouse } from "./types/blueprints/tx";


const types = [
  ["/rafaelsousa.realestate.blueprints.MsgDeleteHouse", MsgDeleteHouse],
  ["/rafaelsousa.realestate.blueprints.MsgCreateHouse", MsgCreateHouse],
  ["/rafaelsousa.realestate.blueprints.MsgUpdateHouse", MsgUpdateHouse],
  
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
    msgDeleteHouse: (data: MsgDeleteHouse): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgDeleteHouse", value: MsgDeleteHouse.fromPartial( data ) }),
    msgCreateHouse: (data: MsgCreateHouse): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgCreateHouse", value: MsgCreateHouse.fromPartial( data ) }),
    msgUpdateHouse: (data: MsgUpdateHouse): EncodeObject => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgUpdateHouse", value: MsgUpdateHouse.fromPartial( data ) }),
    
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
