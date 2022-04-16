import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthLoginRequest {
  email: String
  password: String
}

export default async function (data: AuthLoginRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/login/', data)
}
