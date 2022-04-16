import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthRefreshRequest {
  refreshToken: String
}

export interface AuthRefreshResponse {
  accessToken: string,
  refreshToken: string,
  user: {
    id: number,
    email: string
  }
}

export default async function (data: AuthRefreshRequest): Promise<CustomAxiosResponse<AuthRefreshResponse, AuthRefreshRequest>> {
  return await baseAxios.post('/auth/refresh/', data)
}
