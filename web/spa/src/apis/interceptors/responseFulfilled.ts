import type {AxiosResponse} from "axios";
import type {CustomAxiosResponse} from "@/apis/schemas";

export default function (response: AxiosResponse<any, any>): CustomAxiosResponse {
  console.log('apis interceptors response fulfilled', response)
  return {
    success: true,
    data: response.data
  } as CustomAxiosResponse
}
