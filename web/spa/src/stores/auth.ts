import {defineStore} from "pinia";
import Cookies from 'js-cookie';

const persistRefreshToken = Cookies.get('auth.refreshToken')

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
    accessToken: '',
    refreshToken: '' || persistRefreshToken,
    user: {
      id: 0,
      email: ''
    }
  } as AuthState),
  getters: {
    isAuth(state: AuthState) {
      return state.accessToken !== null
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

      Cookies.set('auth.refreshToken', state.refreshToken)
    }
  }
})
