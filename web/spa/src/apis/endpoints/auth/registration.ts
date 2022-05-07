import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {AuthRegistrationRequest} from "@/apis/schemas/auth";

export default async function (data: AuthRegistrationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/registration/', data)
}
