import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthLoginRequest {
  email: String
  password: String
}

export interface AuthLoginResponse {
  accessToken: string,
  refreshToken: string,
  user: {
    id: number,
    email: string
  }
}

export default async function (data: AuthLoginRequest): Promise<CustomAxiosResponse<AuthLoginResponse, AuthLoginRequest>> {
  return await baseAxios.post('/auth/login/', data)
}
