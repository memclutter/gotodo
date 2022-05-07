import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {AuthLoginRequest, AuthLoginResponse} from "@/apis/schemas/auth";

export default async function (data: AuthLoginRequest): Promise<CustomAxiosResponse<AuthLoginResponse, AuthLoginRequest>> {
  return await baseAxios.post('/auth/login/', data)
}
