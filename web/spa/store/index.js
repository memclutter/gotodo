export const state = () => ({
  auth: {
    accessToken: null,
    refreshToken: {
      token: null
    },
    user: {
      id: null,
      email: null
    }
  },
  login: {
    form: {
      email: null,
      password: null
    }
  },
  tasks: {
    list: {
      items: [],
      totalCount: 0
    }
  }
})

export const actions = {

  async authRefresh({commit, state}) {
    const resp = await this.$axios.post('/auth/refresh/', {refreshToken: state.auth.refreshToken.token})
    console.log(resp)
  },

  async authLogin({commit, state}) {
    const resp = await this.$axios.post('/auth/login/', state.login.form)
  },

  async tasksList({}) {
    const resp = await this.$axios.get('/tasks/')
    console.log(resp)
  }
}
