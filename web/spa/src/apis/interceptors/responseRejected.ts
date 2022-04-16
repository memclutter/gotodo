import {ElMessage} from "element-plus";
import type {AxiosError, AxiosResponse} from "axios";
import authRefresh from '@/apis/endpoints/auth/refresh'
import {useAuthStore} from "@/stores/auth";
import baseAxios from "@/apis/base";

export interface Error {
  message?: String
  validationErrors?: any
}

const validationErrorCodes = {
  'REQUIRED': 'Field is required',
  'EMAIL': 'Invalid email address',
  'USER_EMAIL_EXISTS': 'This email address already use',
}

export default function (error: AxiosError) {
  console.log('apis interceptors response rejected', error)
  try {
    const response: AxiosResponse | undefined = error.response
    const status = response?.status || 0;
    const data: Error = response?.data || new Error()
    const message = data.message || response?.statusText
    const validationErrors = data.validationErrors || {}

    if (status >= 500) {
      ElMessage({
        message: `Server error: ${message}`,
        type: 'error'
      })
    } else if (status === 0) {
      ElMessage({
        message: `Network error: ${error.toString()}`,
        type: 'error'
      })
    } else if (status === 400) {
      for (const field in data.validationErrors) {
        if (Array.isArray(data.validationErrors[field])) {
          data.validationErrors[field] = data.validationErrors[field].map((code: string) => validationErrorCodes[code] || code)
        }
      }
    } else if (status === 401) {
      const authStore = useAuthStore()
      if (authStore.refreshToken) {
        authRefresh({refreshToken: authStore.refreshToken}).then(({success, data}) => {
          if (success) {
            authStore.set(data)
            return baseAxios.request(error.config)
          }
        })
      }
    }

    return {
      success: false,
      message: message,
      validationErrors: validationErrors
    }
  } catch (e) {
    return Promise.reject(e)
  }
}
