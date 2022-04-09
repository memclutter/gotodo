import {ElMessage} from "element-plus";
import type {AxiosResponse} from "axios";

export interface Error {
  message?: String
  validationErrors?: any
}

export default function (error) {
  console.log('apis interceptors response rejected', error)
  try {
    const response: AxiosResponse | object = error.response || {}
    const status = response.status || 0;
    const data: Error = response.data || new Error()
    const message = data.message || response.statusText
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
