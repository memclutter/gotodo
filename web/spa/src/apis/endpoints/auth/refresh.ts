import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthRefreshRequest {
  refreshToken: String
}

export default async function (data: AuthRefreshRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/refresh/', data)
}
