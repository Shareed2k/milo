import _ from 'lodash'

const defaultUser = {
  username: '',
  role: '',
  email: '',
  api_token: '',
  token: '',
  teams: []
}

const state = {
  user: defaultUser,
  isAuthorized: false
}

const mutations = {
  SET_USER (state, user) {
    state.user = user
    state.isAuthorized = true
  },
  UNSET_USER (state) {
    state.user = defaultUser
    state.isAuthorized = false
  },
  SET_TOKEN (state, token) {
    state.token = token
  }
}

const actions = {
  setUser ({ commit }, user) {
    return new Promise((resolve, reject) => {
      if (_.isNil(user)) {
        reject(new Error('user is empty'))
      } else {
        commit('SET_USER', user)
        resolve(user)
      }
    })
  },

  setToken ({ commit }, token) {
    return new Promise(resolve => {
      commit('SET_TOKEN', token)
      resolve(token)
    })
  },

  unsetUser ({ commit }) {
    commit('UNSET_USER')
  }
}

const getters = {
  isAuthorized: state => state.isAuthorized,
  getUser: state => state.user,
  getToken: state => state.user.token,
  getTeamByName: state => name => _.filter(state.user.teams, team => team.name === name),
  getTeamByUuid: state => uuid => _.filter(state.user.teams, team => team.uuid === uuid)
}

export default {
  state,
  mutations,
  actions,
  getters
}
