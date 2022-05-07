import {ElMessage} from "element-plus";
import type {AxiosError, AxiosResponse} from "axios";
import {authRefresh} from "@/apis/endpoints/auth";
import {useAuthStore} from "@/stores/auth";
import {baseAxios} from "@/apis";
import type {CustomAxiosResponse} from "@/apis/schemas";

export interface Error {
  message?: String
  validationErrors?: {[key: string]: string[]}
}

const validationErrorCodes: {[key: string]: string} = {
  'REQUIRED': 'Field is required',
  'EMAIL': 'Invalid email address',
  'USER_EMAIL_EXISTS': 'This email address already use',
}

export default function (error: AxiosError): Promise<CustomAxiosResponse> | CustomAxiosResponse {
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
      for (const field in validationErrors) {
        if (Array.isArray(validationErrors[field])) {
          validationErrors[field] = validationErrors[field].map((code: string) => validationErrorCodes[code] || code)
        }
      }
    } else if (status === 401) {
      const authStore = useAuthStore()
      if (authStore.refreshToken) {
        return authRefresh({refreshToken: authStore.refreshToken}).then(({success, data}) => {
          if (success) {
            authStore.set(data)
            return baseAxios.request(error.config)
          } else {
            // Clear auth store
            authStore.unset()
            return {success: false, message: 'Unauthorized'} as CustomAxiosResponse
          }
        })
      }
    }

    return {
      success: false,
      message: message,
      validationErrors: validationErrors
    } as CustomAxiosResponse
  } catch (e) {
    return Promise.reject(e)
  }
}
