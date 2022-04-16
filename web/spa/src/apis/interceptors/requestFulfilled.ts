import type {AxiosRequestConfig} from "axios";
import {useAuthStore} from "@/stores/auth";

export default function (request: AxiosRequestConfig) {
  console.log('apis interceptors request fulfilled', request)
  const authStore = useAuthStore()
  if ( authStore.isAuth && authStore.accessToken) {
    if (request.headers) {
      request.headers['Authorization'] = `Bearer ${authStore.accessToken}`
    }
  }
  return request
}
