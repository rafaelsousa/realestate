// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteOwner } from "./types/property/tx";
import { MsgCreateProperty } from "./types/property/tx";
import { MsgCreateOwner } from "./types/property/tx";
import { MsgUpdateOwner } from "./types/property/tx";
import { MsgUpdateProperty } from "./types/property/tx";
import { MsgDeleteProperty } from "./types/property/tx";
const types = [
    ["/rafaelsousa.realestate.property.MsgDeleteOwner", MsgDeleteOwner],
    ["/rafaelsousa.realestate.property.MsgCreateProperty", MsgCreateProperty],
    ["/rafaelsousa.realestate.property.MsgCreateOwner", MsgCreateOwner],
    ["/rafaelsousa.realestate.property.MsgUpdateOwner", MsgUpdateOwner],
    ["/rafaelsousa.realestate.property.MsgUpdateProperty", MsgUpdateProperty],
    ["/rafaelsousa.realestate.property.MsgDeleteProperty", MsgDeleteProperty],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgDeleteOwner: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgDeleteOwner", value: data }),
        msgCreateProperty: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgCreateProperty", value: data }),
        msgCreateOwner: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgCreateOwner", value: data }),
        msgUpdateOwner: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgUpdateOwner", value: data }),
        msgUpdateProperty: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgUpdateProperty", value: data }),
        msgDeleteProperty: (data) => ({ typeUrl: "/rafaelsousa.realestate.property.MsgDeleteProperty", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
