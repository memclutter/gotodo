import {defineStore} from "pinia";
import Cookies from 'js-cookie';

const persistRefreshToken = Cookies.get('auth.refreshToken')
const persistAccessToken = Cookies.get('auth.accessToken')

export type AuthStateUser = {
  id: Number
  email: String
}

export type AuthState = {
  accessToken: String
  refreshToken: String
  user: AuthStateUser
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    accessToken: '' || persistAccessToken,
    refreshToken: '' || persistRefreshToken,
    user: {
      id: 0,
      email: ''
    }
  } as AuthState),
  getters: {
    isAuth(state: AuthState) {
      return !!state.accessToken
    },
    userId(state: AuthState) {
      return state.user.id
    },
    userEmail(state: AuthState) {
      return state.user.email
    }
  },
  actions: {
    set(state: AuthState) {
      this.accessToken = state.accessToken
      this.refreshToken = state.refreshToken
      this.user.id = state.user.id
      this.user.email = state.user.email

      Cookies.set('auth.accessToken', state.accessToken)
      Cookies.set('auth.refreshToken', state.refreshToken)
    },
    unset() {
      this.accessToken = ''
      this.refreshToken = ''
      this.user.id = 0
      this.user.email = ''

      Cookies.remove('auth.accessToken')
      Cookies.remove('auth.refreshToken')
    }
  }
})
