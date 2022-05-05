import {defineStore} from "pinia";
import jwtDecode from 'jwt-decode';

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
    accessToken: '',
    refreshToken: '',
    user: {
      id: 0,
      firstName: '',
      lastName: '',
      email: ''
    }
  } as AuthState),
  persist: {
    key: 'auth',
    paths: ['accessToken', 'refreshToken'],
    afterRestore({store}) {
      store.user = store.accessToken ? <AuthStateUser>jwtDecode(<string>store.accessToken) : <AuthStateUser>{}
    }
  },
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
    },
    unset() {
      this.accessToken = ''
      this.refreshToken = ''
      this.user = <AuthStateUser>{}
    }
  }
})
