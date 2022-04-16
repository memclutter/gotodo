import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface AuthRegistrationRequest {
  email: String
  password: String
}

export default async function (data: AuthRegistrationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/registration/', data)
}
