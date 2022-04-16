import type {AxiosResponse} from "axios";

export default function (response: AxiosResponse<any, any>) {
  console.log('apis interceptors response fulfilled', response)
  return {
    success: true,
    data: response.data
  }
}
