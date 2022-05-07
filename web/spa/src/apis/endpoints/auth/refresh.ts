import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {AuthRefreshRequest, AuthRefreshResponse} from "@/apis/schemas/auth";

export default async function (data: AuthRefreshRequest): Promise<CustomAxiosResponse<AuthRefreshResponse, AuthRefreshRequest>> {
  return await baseAxios.post('/auth/refresh/', data)
}
