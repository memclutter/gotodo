import {defineStore} from "pinia";
import jwtDecode from 'jwt-decode';
import type {User} from "@/models/users";

export type AuthState = {
  accessToken: String
  refreshToken: String
  user: User
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
      store.user = store.accessToken ? <User>jwtDecode(<string>store.accessToken) : <User>{}
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
