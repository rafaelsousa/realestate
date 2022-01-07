import { txClient, queryClient, MissingWalletError, registry } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'
import { House } from './module/types/blueprints/house'
import { Params } from './module/types/blueprints/params'
export { House, Params }
async function initTxClient(vuexGetters) {
  return await txClient(vuexGetters['common/wallet/signer'], {
    addr: vuexGetters['common/env/apiTendermint'],
  })
}
async function initQueryClient(vuexGetters) {
  return await queryClient({
    addr: vuexGetters['common/env/apiCosmos'],
  })
}
function mergeResults(value, next_values) {
  for (let prop of Object.keys(next_values)) {
    if (Array.isArray(next_values[prop])) {
      value[prop] = [...value[prop], ...next_values[prop]]
    } else {
      value[prop] = next_values[prop]
    }
  }
  return value
}
function getStructure(template) {
  let structure = { fields: [] }
  for (const [key, value] of Object.entries(template)) {
    let field = {}
    field.name = key
    field.type = typeof value
    structure.fields.push(field)
  }
  return structure
}
const getDefaultState = () => {
  return {
    Params: {},
    House: {},
    HouseAll: {},
    _Structure: {
      House: getStructure(House.fromPartial({})),
      Params: getStructure(Params.fromPartial({})),
    },
    _Registry: registry,
    _Subscriptions: new Set(),
  }
}
// initial state
const state = getDefaultState()
export default {
  namespaced: true,
  state,
  mutations: {
    RESET_STATE(state) {
      Object.assign(state, getDefaultState())
    },
    QUERY(state, { query, key, value }) {
      state[query][JSON.stringify(key)] = value
    },
    SUBSCRIBE(state, subscription) {
      state._Subscriptions.add(JSON.stringify(subscription))
    },
    UNSUBSCRIBE(state, subscription) {
      state._Subscriptions.delete(JSON.stringify(subscription))
    },
  },
  getters: {
    getParams:
      (state) =>
      (params = { params: {} }) => {
        if (!params.query) {
          params.query = null
        }
        return state.Params[JSON.stringify(params)] ?? {}
      },
    getHouse:
      (state) =>
      (params = { params: {} }) => {
        if (!params.query) {
          params.query = null
        }
        return state.House[JSON.stringify(params)] ?? {}
      },
    getHouseAll:
      (state) =>
      (params = { params: {} }) => {
        if (!params.query) {
          params.query = null
        }
        return state.HouseAll[JSON.stringify(params)] ?? {}
      },
    getTypeStructure: (state) => (type) => {
      return state._Structure[type].fields
    },
    getRegistry: (state) => {
      return state._Registry
    },
  },
  actions: {
    init({ dispatch, rootGetters }) {
      console.log('Vuex module: rafaelsousa.realestate.blueprints initialized!')
      if (rootGetters['common/env/client']) {
        rootGetters['common/env/client'].on('newblock', () => {
          dispatch('StoreUpdate')
        })
      }
    },
    resetState({ commit }) {
      commit('RESET_STATE')
    },
    unsubscribe({ commit }, subscription) {
      commit('UNSUBSCRIBE', subscription)
    },
    async StoreUpdate({ state, dispatch }) {
      state._Subscriptions.forEach(async (subscription) => {
        try {
          const sub = JSON.parse(subscription)
          await dispatch(sub.action, sub.payload)
        } catch (e) {
          throw new SpVuexError('Subscriptions: ' + e.message)
        }
      })
    },
    async QueryParams(
      { commit, rootGetters, getters },
      { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null },
    ) {
      try {
        const key = params ?? {}
        const queryClient = await initQueryClient(rootGetters)
        let value = (await queryClient.queryParams()).data
        commit('QUERY', { query: 'Params', key: { params: { ...key }, query }, value })
        if (subscribe)
          commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: { ...key }, query } })
        return getters['getParams']({ params: { ...key }, query }) ?? {}
      } catch (e) {
        throw new SpVuexError('QueryClient:QueryParams', 'API Node Unavailable. Could not perform query: ' + e.message)
      }
    },
    async QueryHouse(
      { commit, rootGetters, getters },
      { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null },
    ) {
      try {
        const key = params ?? {}
        const queryClient = await initQueryClient(rootGetters)
        let value = (await queryClient.queryHouse(key.id)).data
        commit('QUERY', { query: 'House', key: { params: { ...key }, query }, value })
        if (subscribe)
          commit('SUBSCRIBE', { action: 'QueryHouse', payload: { options: { all }, params: { ...key }, query } })
        return getters['getHouse']({ params: { ...key }, query }) ?? {}
      } catch (e) {
        throw new SpVuexError('QueryClient:QueryHouse', 'API Node Unavailable. Could not perform query: ' + e.message)
      }
    },
    async QueryHouseAll(
      { commit, rootGetters, getters },
      { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null },
    ) {
      try {
        const key = params ?? {}
        const queryClient = await initQueryClient(rootGetters)
        let value = (await queryClient.queryHouseAll(query)).data
        while (all && value.pagination && value.pagination.next_key != null) {
          let next_values = (await queryClient.queryHouseAll({ ...query, 'pagination.key': value.pagination.next_key }))
            .data
          value = mergeResults(value, next_values)
        }
        commit('QUERY', { query: 'HouseAll', key: { params: { ...key }, query }, value })
        if (subscribe)
          commit('SUBSCRIBE', { action: 'QueryHouseAll', payload: { options: { all }, params: { ...key }, query } })
        return getters['getHouseAll']({ params: { ...key }, query }) ?? {}
      } catch (e) {
        throw new SpVuexError(
          'QueryClient:QueryHouseAll',
          'API Node Unavailable. Could not perform query: ' + e.message,
        )
      }
    },
    async sendMsgCreateHouse({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgCreateHouse(value)
        const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee, gas: '200000' }, memo })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgCreateHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgCreateHouse:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async sendMsgUpdateHouse({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgUpdateHouse(value)
        const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee, gas: '200000' }, memo })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgUpdateHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgUpdateHouse:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async sendMsgDeleteHouse({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgDeleteHouse(value)
        const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee, gas: '200000' }, memo })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgDeleteHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgDeleteHouse:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async MsgCreateHouse({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgCreateHouse(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgCreateHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgCreateHouse:Create', 'Could not create message: ' + e.message)
        }
      }
    },
    async MsgUpdateHouse({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgUpdateHouse(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgUpdateHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgUpdateHouse:Create', 'Could not create message: ' + e.message)
        }
      }
    },
    async MsgDeleteHouse({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgDeleteHouse(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgDeleteHouse:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgDeleteHouse:Create', 'Could not create message: ' + e.message)
        }
      }
    },
  },
}
