import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthConfirmationRequest {
  token: String
}

export default async function (data: AuthConfirmationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/confirmation/', data)
}
