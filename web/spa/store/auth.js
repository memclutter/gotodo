import Vue from "vue";

export const state = {
  session: {
    user: {
      id: null,
      email: null
    },
    accessToken: null,
    refreshToken: null
  },

}

export const mutations = {
  setSession: (state, session) => Vue.set(state, 'session', session),
}

export const actions = {
  async login({commit, state}, reqData) {
    const resp = await this.$axios.post('/auth/login/', reqData);
    if (resp.status === 400) {
      return {
        email: 'Invalid email or password'
      }
    }

    commit('setSession', resp.data)
    return {}
  },
  async registration({commit, state}, reqData) {
    const resp = await this.$axios.post('/auth/registration/', reqData);
    if (resp.hasOwnProperty('success') && resp.success) {
      return {}
    }
    return resp.data;
  }
}
