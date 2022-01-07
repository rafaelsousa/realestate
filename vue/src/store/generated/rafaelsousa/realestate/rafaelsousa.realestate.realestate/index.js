import { MissingWalletError, queryClient, registry, txClient } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'
import { Params } from './module/types/realestate/params'
import { Property } from './module/types/realestate/property'

export { Params, Property }

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
    Property: {},
    PropertyAll: {},
    _Structure: {
      Params: getStructure(Params.fromPartial({})),
      Property: getStructure(Property.fromPartial({})),
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
    getProperty:
      (state) =>
      (params = { params: {} }) => {
        if (!params.query) {
          params.query = null
        }
        return state.Property[JSON.stringify(params)] ?? {}
      },
    getPropertyAll:
      (state) =>
      (params = { params: {} }) => {
        if (!params.query) {
          params.query = null
        }
        return state.PropertyAll[JSON.stringify(params)] ?? {}
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
      console.log('Vuex module: rafaelsousa.realestate.realestate initialized!')
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
    async QueryProperty(
      { commit, rootGetters, getters },
      { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null },
    ) {
      try {
        const key = params ?? {}
        const queryClient = await initQueryClient(rootGetters)
        let value = (await queryClient.queryProperty(key.id)).data
        commit('QUERY', { query: 'Property', key: { params: { ...key }, query }, value })
        if (subscribe)
          commit('SUBSCRIBE', { action: 'QueryProperty', payload: { options: { all }, params: { ...key }, query } })
        return getters['getProperty']({ params: { ...key }, query }) ?? {}
      } catch (e) {
        throw new SpVuexError(
          'QueryClient:QueryProperty',
          'API Node Unavailable. Could not perform query: ' + e.message,
        )
      }
    },
    async QueryPropertyAll(
      { commit, rootGetters, getters },
      { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null },
    ) {
      try {
        const key = params ?? {}
        const queryClient = await initQueryClient(rootGetters)
        let value = (await queryClient.queryPropertyAll(query)).data
        while (all && value.pagination && value.pagination.next_key != null) {
          let next_values = (
            await queryClient.queryPropertyAll({ ...query, 'pagination.key': value.pagination.next_key })
          ).data
          value = mergeResults(value, next_values)
        }
        commit('QUERY', { query: 'PropertyAll', key: { params: { ...key }, query }, value })
        if (subscribe)
          commit('SUBSCRIBE', { action: 'QueryPropertyAll', payload: { options: { all }, params: { ...key }, query } })
        return getters['getPropertyAll']({ params: { ...key }, query }) ?? {}
      } catch (e) {
        throw new SpVuexError(
          'QueryClient:QueryPropertyAll',
          'API Node Unavailable. Could not perform query: ' + e.message,
        )
      }
    },
    async sendMsgDeleteProperty({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgDeleteProperty(value)
        const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee, gas: '200000' }, memo })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgDeleteProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgDeleteProperty:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async sendMsgCreateProperty({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgCreateProperty(value)
        const result = await txClient.signAndBroadcast([msg], {
          fee: {
            amount: fee,
            gas: '200000',
          },
          memo,
        })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgCreateProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgCreateProperty:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async sendMsgUpdateProperty({ rootGetters }, { value, fee = [], memo = '' }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgUpdateProperty(value)
        const result = await txClient.signAndBroadcast([msg], {
          fee: {
            amount: fee,
            gas: '200000',
          },
          memo,
        })
        return result
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgUpdateProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgUpdateProperty:Send', 'Could not broadcast Tx: ' + e.message)
        }
      }
    },
    async MsgDeleteProperty({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgDeleteProperty(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgDeleteProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgDeleteProperty:Create', 'Could not create message: ' + e.message)
        }
      }
    },
    async MsgCreateProperty({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgCreateProperty(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgCreateProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgCreateProperty:Create', 'Could not create message: ' + e.message)
        }
      }
    },
    async MsgUpdateProperty({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters)
        const msg = await txClient.msgUpdateProperty(value)
        return msg
      } catch (e) {
        if (e == MissingWalletError) {
          throw new SpVuexError(
            'TxClient:MsgUpdateProperty:Init',
            'Could not initialize signing client. Wallet is required.',
          )
        } else {
          throw new SpVuexError('TxClient:MsgUpdateProperty:Create', 'Could not create message: ' + e.message)
        }
      }
    },
  },
}
