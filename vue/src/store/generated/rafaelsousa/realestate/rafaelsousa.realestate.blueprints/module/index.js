// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateHouse } from "./types/blueprints/tx";
import { MsgUpdateHouse } from "./types/blueprints/tx";
import { MsgDeleteHouse } from "./types/blueprints/tx";
const types = [
    ["/rafaelsousa.realestate.blueprints.MsgCreateHouse", MsgCreateHouse],
    ["/rafaelsousa.realestate.blueprints.MsgUpdateHouse", MsgUpdateHouse],
    ["/rafaelsousa.realestate.blueprints.MsgDeleteHouse", MsgDeleteHouse],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreateHouse: (data) => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgCreateHouse", value: MsgCreateHouse.fromPartial(data) }),
        msgUpdateHouse: (data) => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgUpdateHouse", value: MsgUpdateHouse.fromPartial(data) }),
        msgDeleteHouse: (data) => ({ typeUrl: "/rafaelsousa.realestate.blueprints.MsgDeleteHouse", value: MsgDeleteHouse.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
