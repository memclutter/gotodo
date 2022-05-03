import {defineStore} from "pinia";
import Cookies from 'js-cookie';
import jwtDecode from 'jwt-decode';

const persistRefreshToken = Cookies.get('auth.refreshToken')
const persistAccessToken = Cookies.get('auth.accessToken')
const persistUser = persistAccessToken ? <AuthStateUser>jwtDecode(<string>persistAccessToken) : <AuthStateUser>{}

export type AuthStateUser = {
  id: Number,
  firstName: String,
  lastName: String,
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
      firstName: '',
      lastName: '',
      email: ''
    } || persistUser
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
      this.user.firstName = state.user.firstName
      this.user.lastName = state.user.lastName
      this.user.email = state.user.email

      Cookies.set('auth.accessToken', state.accessToken)
      Cookies.set('auth.refreshToken', state.refreshToken)
    },
    unset() {
      this.accessToken = ''
      this.refreshToken = ''
      this.user = <AuthStateUser>{}

      Cookies.remove('auth.accessToken')
      Cookies.remove('auth.refreshToken')
    }
  }
})
